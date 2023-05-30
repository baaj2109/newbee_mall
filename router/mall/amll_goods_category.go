package mall

import (
	v1 "github.com/baaj2109/newbee_mall/api/v1"
	"github.com/gin-gonic/gin"
)

type MallGoodsCategoryIndexRouter struct {
}

func (m *MallGoodsInfoIndexRouter) InitMallGoodsCategoryIndexRouter(Router *gin.RouterGroup) {
	mallGoodsRouter := Router.Group("v1")
	var mallGoodsCategoryApi = v1.ApiGroupApp.MallApiGroup.MallGoodsCategoryApi
	{
		mallGoodsRouter.GET("categories", mallGoodsCategoryApi.GetGoodsCategory) // 获取分类数据
	}
}
