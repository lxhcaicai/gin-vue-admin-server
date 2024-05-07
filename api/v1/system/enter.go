package system

import "github.com/lxhcaicai/gin-vue-admin/server/service"

type ApiGroup struct {
	BaseApi
	DBApi
	SystemApi
	JwtApi
	OperationRecordApi
	AuthorityApi
	CasbinApi
	SystemApiApi
	DictionaryDetailApi
	DictionaryApi
	AuthorityBtnApi
	AuthorityMenuApi
	SysExportTemplateApi
}

var (
	userService             = service.ServiceGroupApp.SystemServiceGroup.UserService
	jwtService              = service.ServiceGroupApp.SystemServiceGroup.JwtService
	initDBService           = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	systemConfigService     = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
	operationRecordService  = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
	authorityService        = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	casbinService           = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	apiService              = service.ServiceGroupApp.SystemServiceGroup.ApiService
	dictionaryDetailService = service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService
	dictionaryService       = service.ServiceGroupApp.SystemServiceGroup.DictionaryService
	authorityBtnService     = service.ServiceGroupApp.SystemServiceGroup.AuthorityBtnService
	menuService             = service.ServiceGroupApp.SystemServiceGroup.MenuService
	baseMenuService         = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
)
