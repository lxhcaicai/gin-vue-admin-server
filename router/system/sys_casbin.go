package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lxhcaicai/gin-vue-admin/server/api/v1"
	"github.com/lxhcaicai/gin-vue-admin/server/middleware"
)

type CasbinRouter struct {
}

func (s *CasbinRouter) InitCasbinRouter(Router *gin.RouterGroup) {
	casbinRouterWithoutRecord := Router.Group("casbin")
	casbinRouter := Router.Group("casbin").Use(middleware.OperationRecord())
	casbinApi := v1.ApiGroupApp.SystemApiGroup.CasbinApi
	{
		casbinRouterWithoutRecord.POST("getPolicyPathByAuthorityId", casbinApi.GetPolicyPathByAuthorityId)
	}
	{
		casbinRouter.POST("updateCasbin", casbinApi.UpdateCasbin)
	}
}
