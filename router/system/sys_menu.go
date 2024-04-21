package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lxhcaicai/gin-vue-admin/server/api/v1"
	"github.com/lxhcaicai/gin-vue-admin/server/middleware"
)

type MenuRouter struct {
}

func (m *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	menuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.AuthorityMenuApi
	{
		menuRouter.POST("addBaseMenu", authorityMenuApi.AddBaseMenu)
	}

}
