package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lxhcaicai/gin-vue-admin/server/api/v1"
)

type UserRouter struct {
}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouterWithoutRecord := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouterWithoutRecord.GET("getUserInfo", baseApi.GetUserInfo) // 获取用户信息
	}
}
