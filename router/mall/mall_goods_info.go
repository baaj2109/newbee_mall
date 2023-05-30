package mall

import (
	v1 "github.com/baaj2109/newbee_mall/api/v1"
	"github.com/gin-gonic/gin"
)

type MallGoodsInfoIndexRouter struct {
}

func (m *MallGoodsInfoIndexRouter) InitMallGoodsInfoIndexRouter(Router *gin.RouterGroup) {
	mallGoodsRouter := Router.Group("v1")
	var mallGoodsInfoApi = v1.ApiGroupApp.MallApiGroup.MallGoodsInfoApi
	{
		mallGoodsRouter.GET("/search", mallGoodsInfoApi.GoodsSearch)           // 商品搜索
		mallGoodsRouter.GET("/goods/detail/:id", mallGoodsInfoApi.GoodsDetail) //商品詳情
	}
}
