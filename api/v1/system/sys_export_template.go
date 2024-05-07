package system

import (
	"github.com/gin-gonic/gin"
	"github.com/lxhcaicai/gin-vue-admin/server/global"
	"github.com/lxhcaicai/gin-vue-admin/server/model/common/response"
	"github.com/lxhcaicai/gin-vue-admin/server/model/system"
	"github.com/lxhcaicai/gin-vue-admin/server/service"
	"github.com/lxhcaicai/gin-vue-admin/server/utils"
	"go.uber.org/zap"
)

type SysExportTemplateApi struct {
}

var sysExportTemplateService = service.ServiceGroupApp.SystemServiceGroup.SysExportTemplateService

// CreateSysExportTemplate 创建导出模板
// @Tags SysExportTemplate
// @Summary 创建导出模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysExportTemplate true "创建导出模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysExportTemplate/createSysExportTemplate [post]
func (sysExportTemplateApi *SysExportTemplateApi) CreateSysExportTemplate(c *gin.Context) {
	var sysExportTemplate system.SysExportTemplate
	err := c.ShouldBindJSON(&sysExportTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name": {
			utils.NotEmpty(),
		},
	}
	if err := utils.Verify(sysExportTemplate, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysExportTemplateService.CreateSysExportTemplate(&sysExportTemplate); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
