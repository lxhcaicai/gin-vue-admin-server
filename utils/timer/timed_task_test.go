package timer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func mockFunc() {
	time.Sleep(time.Second)
	fmt.Println("1s...")
}

func TestNewTimerTask(t *testing.T) {
	tm := NewTimerTask()
	_tm := tm.(*timer)

	{
		_, err := tm.AddTaskByFunc("func", "@every 1s", mockFunc, "测试mockfunc")
		assert.Nil(t, err)
		_, ok := _tm.cronList["func"]
		if !ok {
			t.Error("no find func")
		}
	}
}
