package system

import "github.com/lxhcaicai/gin-vue-admin/server/service"

type ApiGroup struct {
	BaseApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
	jwtService  = service.ServiceGroupApp.SystemServiceGroup.JwtService
)
