package mall

import (
	v1 "github.com/baaj2109/newbee_mall/api/v1"
	"github.com/baaj2109/newbee_mall/middleware"
	"github.com/gin-gonic/gin"
)

type MallOrderRouter struct {
}

func (m *MallOrderRouter) InitMallOrderRouter(Router *gin.RouterGroup) {
	mallOrderRouter := Router.Group("v1").Use(middleware.UserJWTAuth())

	var mallOrderRouterApi = v1.ApiGroupApp.MallApiGroup.MallOrderApi
	{
		mallOrderRouter.GET("/paySuccess", mallOrderRouterApi.PaySuccess)             //模擬支付成功回調的接口
		mallOrderRouter.PUT("/order/:orderNo/finish", mallOrderRouterApi.FinishOrder) //確認收貨接口
		mallOrderRouter.PUT("/order/:orderNo/cancel", mallOrderRouterApi.CancelOrder) //取消訂單接口
		mallOrderRouter.GET("/order/:orderNo", mallOrderRouterApi.OrderDetailPage)    //訂單詳情接口
		mallOrderRouter.GET("/order", mallOrderRouterApi.OrderList)                   //訂單列表接口
		mallOrderRouter.POST("/saveOrder", mallOrderRouterApi.SaveOrder)

	}
}
