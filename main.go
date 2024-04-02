package main

import (
	"github.com/lxhcaicai/gin-vue-admin/server/core"
	"github.com/lxhcaicai/gin-vue-admin/server/global"
	"github.com/lxhcaicai/gin-vue-admin/server/initialize"
	"go.uber.org/zap"
)

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap()
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() //gorm连接数据库
	initialize.Timer()                //定时任务

	core.RunWindowsServer()
}
