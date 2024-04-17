package system

import (
	"github.com/gin-gonic/gin"
	"github.com/lxhcaicai/gin-vue-admin/server/global"
	"github.com/lxhcaicai/gin-vue-admin/server/model/common/response"
	"github.com/lxhcaicai/gin-vue-admin/server/model/system"
	"go.uber.org/zap"
)

type DictionaryDetailApi struct {
}

// CreateSysDictionaryDetail
// @Tags      SysDictionaryDetail
// @Summary   创建SysDictionaryDetail
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysDictionaryDetail     true  "SysDictionaryDetail模型"
// @Success   200   {object}  response.Response{msg=string}  "创建SysDictionaryDetail"
// @Router    /sysDictionaryDetail/createSysDictionaryDetail [post]
func (d *DictionaryDetailApi) CreateSysDictionaryDetail(c *gin.Context) {
	var detail system.SysDictionaryDetail
	err := c.ShouldBindJSON(&detail)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryDetailService.CreateSysDictionaryDetail(detail)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSysDictionaryDetail
// @Tags      SysDictionaryDetail
// @Summary   删除SysDictionaryDetail
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysDictionaryDetail     true  "SysDictionaryDetail模型"
// @Success   200   {object}  response.Response{msg=string}  "删除SysDictionaryDetail"
// @Router    /sysDictionaryDetail/deleteSysDictionaryDetail [delete]
func (d DictionaryDetailApi) DeleteSysDictionaryDetail(c *gin.Context) {
	var detail system.SysDictionaryDetail
	err := c.ShouldBindJSON(&detail)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryDetailService.DeleteSysDictionaryDetail(detail)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
