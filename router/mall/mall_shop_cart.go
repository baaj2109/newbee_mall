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
		mallShopCartRouter.GET("/shop-cart", mallShopCartApi.GetCartItemList)                                 //购物车列表(网页移动端不分页)
		mallShopCartRouter.POST("/shop-cart", mallShopCartApi.SaveCartItem)                                   //添加购物车
		mallShopCartRouter.PUT("/shop-cart", mallShopCartApi.UpdateCartItem)                                  //修改购物车
		mallShopCartRouter.DELETE("/shop-cart/:newBeeMallShoppingCartItemId", mallShopCartApi.DeleteCartItem) //删除购物车
		mallShopCartRouter.GET("/shop-cart/settle", mallShopCartApi.ToSettle)                                 //根据购物项id数组查询购物项明细

	}
}
