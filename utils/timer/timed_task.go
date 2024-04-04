package timer

import (
	"github.com/robfig/cron/v3"
	"sync"
)

type Timer interface {

	// 通过函数的方法添加任务
	AddTaskByFunc(cronName string, spec string, task func(), taskName string, option ...cron.Option) (cron.EntryID, error)
}

type task struct {
	EntryID  cron.EntryID
	Spec     string
	TaskName string
}

type taskManager struct {
	cron  *cron.Cron
	tasks map[cron.EntryID]*task
}

type timer struct {
	cronList map[string]*taskManager
	sync.Mutex
}

func (t *timer) AddTaskByFunc(cronName string, spec string, fun func(), taskName string, option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	//if _, ok := t.cronList[cronName]; ok {
	//	tasks := make(map[cron.EntryID]*task)
	//	t.cronList[cronName] = &taskManager{
	//		cron:  cron.New(option...),
	//		tasks: tasks,
	//	}
	//}
	//id, err := t.cronList[cronName].cron.AddFunc(spec, fun)
	//t.cronList[cronName].cron.Start()
	//t.cronList[cronName].tasks[id] = &task{
	//	EntryID:  id,
	//	Spec:     spec,
	//	TaskName: taskName,
	//}
	//return id, err
	if _, ok := t.cronList[cronName]; !ok {
		tasks := make(map[cron.EntryID]*task)
		t.cronList[cronName] = &taskManager{
			cron:  cron.New(option...),
			tasks: tasks,
		}
	}
	id, err := t.cronList[cronName].cron.AddFunc(spec, fun)
	t.cronList[cronName].cron.Start()
	t.cronList[cronName].tasks[id] = &task{
		EntryID:  id,
		Spec:     spec,
		TaskName: taskName,
	}
	return id, err
}

func NewTimerTask() Timer {
	return &timer{cronList: make(map[string]*taskManager)}
}
