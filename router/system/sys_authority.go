package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lxhcaicai/gin-vue-admin/server/api/v1"
	"github.com/lxhcaicai/gin-vue-admin/server/middleware"
)

type AuthorityRouter struct {
}

func (s *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	authorityRouterWithoutRecord := Router.Group("authority")
	authorityRouter := Router.Group("authority").Use(middleware.OperationRecord())
	authorityApi := v1.ApiGroupApp.SystemApiGroup.AuthorityApi
	{
		authorityRouterWithoutRecord.POST("getAuthorityList", authorityApi.GetAuthorityList)
	}
	{
		authorityRouter.POST("createAuthority", authorityApi.CreateAuthority) // 创建角色
		authorityRouter.POST("deleteAuthority", authorityApi.DeleteAuthority) // 删除角色
		authorityRouter.POST("updateAuthority", authorityApi.UpdateAuthority)
	}
}
