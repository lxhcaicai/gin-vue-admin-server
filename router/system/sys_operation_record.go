package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lxhcaicai/gin-vue-admin/server/api/v1"
)

type OperationRecordRouter struct {
}

func (s *OperationRecordRouter) InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	operationRecordRouter := Router.Group("sysOperationRecord")
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.OperationRecordApi
	{
		operationRecordRouter.GET("getSysOperationRecordList", authorityMenuApi.GetSysOperationRecordList)
		operationRecordRouter.GET("findSysOperationRecord", authorityMenuApi.FindSysOperationRecord)
	}
}
