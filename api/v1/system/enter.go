package system

import "github.com/lxhcaicai/gin-vue-admin/server/service"

type ApiGroup struct {
	BaseApi
	DBApi
	SystemApi
}

var (
	userService   = service.ServiceGroupApp.SystemServiceGroup.UserService
	jwtService    = service.ServiceGroupApp.SystemServiceGroup.JwtService
	initDBService = service.ServiceGroupApp.SystemServiceGroup.InitDBService
)
