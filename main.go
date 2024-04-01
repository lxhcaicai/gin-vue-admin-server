package main

import (
	"github.com/lxhcaicai/gin-vue-admin/server/core"
	"github.com/lxhcaicai/gin-vue-admin/server/global"
)

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
}
