package router

import (
	"net/http"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/middleware"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 為用戶頭像和文件提供靜態地址
	//Router.Use(middleware.LoadTls())  // 打開就能玩https了
	global.GVA_LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打開

	PublicGroup := Router.Group("")
	{
		// 健康監測
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{

		global.GVA_LOG.Info("use middleware cors")
		// 方便统一添加路由组前缀 多服務器上線使用
		//商城後台路由
		manageRouter := RouterGroupApp.Manage
		ManageGroup := Router.Group("manage-api")

		//商城後台路由初始化
		manageRouter.InitManageAdminUserRouter(ManageGroup)
		manageRouter.InitManageGoodsCategoryRouter(ManageGroup)
		manageRouter.InitManageGoodsInfoRouter(ManageGroup)
		manageRouter.InitManageCarouselRouter(ManageGroup)
		manageRouter.InitManageIndexConfigRouter(ManageGroup)
		manageRouter.InitManageOrderRouter(ManageGroup)
	}
	//mall frount end
	mallRouter := RouterGroupApp.Mall
	MallGroup := Router.Group("api")
	{
		// mall frount end
		// mallRouter.InitMallCarouselIndexRouter(MallGroup)
		mallRouter.InitMallGoodsInfoIndexRouter(MallGroup)
		mallRouter.InitMallGoodsCategoryIndexRouter(MallGroup)
		mallRouter.InitMallUserRouter(MallGroup)
		mallRouter.InitMallUserAddressRouter(MallGroup)
		mallRouter.InitMallShopCartRouter(MallGroup)
		mallRouter.InitMallOrderRouter(MallGroup)
	}
	global.GVA_LOG.Info("router register success")
	return Router
}
