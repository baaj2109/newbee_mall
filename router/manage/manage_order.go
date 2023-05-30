package manage

import (
	v1 "github.com/baaj2109/newbee_mall/api/v1"
	"github.com/baaj2109/newbee_mall/middleware"
	"github.com/gin-gonic/gin"
)

type ManageOrderRouter struct {
}

func (r *ManageOrderRouter) InitManageOrderRouter(Router *gin.RouterGroup) {
	mallOrderRouter := Router.Group("v1").Use(middleware.AdminJWTAuth())
	var mallOrderApi = v1.ApiGroupApp.ManageApiGroup.ManageOrderApi
	{
		mallOrderRouter.PUT("orders/checkDone", mallOrderApi.CheckDoneOrder) // 發貨
		mallOrderRouter.PUT("orders/checkOut", mallOrderApi.CheckOutOrder)   // 出庫
		mallOrderRouter.PUT("orders/close", mallOrderApi.CloseOrder)         // 出庫
		// mallOrderRouter.GET("orders/:orderId", mallOrderApi.FindMallOrder)   // 根據ID獲取MallOrder
		// mallOrderRouter.GET("orders", mallOrderApi.GetMallOrderList)         // 獲取MallOrder列表
	}
}
