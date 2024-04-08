package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lxhcaicai/gin-vue-admin/server/api/v1"
	"github.com/lxhcaicai/gin-vue-admin/server/middleware"
)

type UserRouter struct {
}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouterWithoutRecord := Router.Group("user")
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList) // 分页获取用户列表
		userRouterWithoutRecord.GET("getUserInfo", baseApi.GetUserInfo)  // 获取用户信息
	}
	{
		userRouter.POST("admin_register", baseApi.Register)       // 管理员注册账号
		userRouter.POST("changePassword", baseApi.ChangePassword) // 用户修改密码
	}
}
