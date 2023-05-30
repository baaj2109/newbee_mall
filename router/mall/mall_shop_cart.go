package mall

import (
	v1 "github.com/baaj2109/newbee_mall/api/v1"
	"github.com/baaj2109/newbee_mall/middleware"
	"github.com/gin-gonic/gin"
)

type MallShopCartRouter struct {
}

func (m *MallUserRouter) InitMallShopCartRouter(Router *gin.RouterGroup) {
	mallShopCartRouter := Router.Group("v1").Use(middleware.UserJWTAuth())
	var mallShopCartApi = v1.ApiGroupApp.MallApiGroup.MallShopCartApi
	{
		mallShopCartRouter.GET("/shop-cart", mallShopCartApi.GetCartItemList)                                 //購物車列表(網頁移動端不分頁)
		mallShopCartRouter.POST("/shop-cart", mallShopCartApi.SaveCartItem)                                   //添加購物車
		mallShopCartRouter.PUT("/shop-cart", mallShopCartApi.UpdateCartItem)                                  //修改購物車
		mallShopCartRouter.DELETE("/shop-cart/:newBeeMallShoppingCartItemId", mallShopCartApi.DeleteCartItem) //刪除購物車
		mallShopCartRouter.GET("/shop-cart/settle", mallShopCartApi.ToSettle)                                 //根據購物項id數組查詢購物項明细

	}
}
