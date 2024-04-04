package router

import "github.com/lxhcaicai/gin-vue-admin/server/router/system"

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
