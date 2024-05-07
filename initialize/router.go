package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/lxhcaicai/gin-vue-admin/server/docs"
	"github.com/lxhcaicai/gin-vue-admin/server/global"
	"github.com/lxhcaicai/gin-vue-admin/server/middleware"
	"github.com/lxhcaicai/gin-vue-admin/server/router"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example

	docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	Router.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) //注册基础功能路由
		systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}

	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitSystemRouter(PrivateGroup)                 //system相关路由
		systemRouter.InitUserRouter(PrivateGroup)                   //注册用户路由
		systemRouter.InitJwtRouter(PrivateGroup)                    //jwt相关路由
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)     // 操作记录
		systemRouter.InitAuthorityRouter(PrivateGroup)              // 注册角色路由
		systemRouter.InitCasbinRouter(PrivateGroup)                 // 权限相关路由
		systemRouter.InitApiRouter(PrivateGroup, PublicGroup)       // 注册功能api路由
		systemRouter.InitSysDictionaryRouter(PrivateGroup)          // 字典管理
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)    // 字典详情管理
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由
		exampleRouter.InitCustomerRouter(PrivateGroup)              // 客户路由
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)     // 权限按钮管理
		systemRouter.InitMenuRouter(PrivateGroup)                   // 注册menu路由
		systemRouter.InitSysExportTemplateRouter(PrivateGroup)
	}
	global.GVA_LOG.Info("router register success")
	return Router
}
