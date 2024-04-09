package system

import "github.com/lxhcaicai/gin-vue-admin/server/service"

type ApiGroup struct {
	BaseApi
	DBApi
	SystemApi
	JwtApi
	OperationRecordApi
}

var (
	userService            = service.ServiceGroupApp.SystemServiceGroup.UserService
	jwtService             = service.ServiceGroupApp.SystemServiceGroup.JwtService
	initDBService          = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	systemConfigService    = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
	operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
)
