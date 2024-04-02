package initialize

import (
	"fmt"
	"github.com/lxhcaicai/gin-vue-admin/server/global"
	"github.com/robfig/cron/v3"
)

func Timer() {
	go func() {
		var option []cron.Option
		option = append(option, cron.WithSeconds())
		// 定时任务demo
		_, err := global.GVA_Timer.AddTaskByFunc("TestTimer", "@every 5s", func() {
			fmt.Println("正在做定时任务...")
		}, "定时做任务", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}
	}()
}
