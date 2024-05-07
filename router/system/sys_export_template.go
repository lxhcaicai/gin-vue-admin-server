package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lxhcaicai/gin-vue-admin/server/api/v1"
	"github.com/lxhcaicai/gin-vue-admin/server/middleware"
)

type SysExportTemplateRouter struct {
}

// InitSysExportTemplateRouter 初始化 导出模板 路由信息
func (s *SysExportTemplateRouter) InitSysExportTemplateRouter(Router *gin.RouterGroup) {
	sysExportTemplateRouter := Router.Group("syExportTemplate").Use(middleware.OperationRecord())
	var sysExportTemplateApi = v1.ApiGroupApp.SystemApiGroup.SysExportTemplateApi
	{
		sysExportTemplateRouter.POST("createSysExportTemplate", sysExportTemplateApi.CreateSysExportTemplate)
		sysExportTemplateRouter.DELETE("deleteSysExportTemplate", sysExportTemplateApi.DeleteSysExportTemplate)
		sysExportTemplateRouter.DELETE("deleteSysExportTemplateByIds", sysExportTemplateApi.DeleteSysExportTemplateByIds)
		sysExportTemplateRouter.PUT("updateSysExportTemplate", sysExportTemplateApi.UpdateSysExportTemplate)
	}
}
