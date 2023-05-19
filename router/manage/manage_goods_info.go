package manage

import (
	v1 "github.com/baaj2109/newbee_mall/api/v1"
	"github.com/gin-gonic/gin"
)

type ManageGoodsInfoRouter struct {
}

func (m *ManageGoodsInfoRouter) InitManageGoodsInfoRouter(Router *gin.RouterGroup) {
	mallGoodsInfoRouter := Router.Group("v1")
	var mallGoodsInfoApi = v1.ApiGroupApp.ManageApiGroup.ManageGoodsInfoApi
	{
		mallGoodsInfoRouter.POST("goods", mallGoodsInfoApi.CreateGoodsInfo)                 // 新建MallGoodsInfo
		mallGoodsInfoRouter.DELETE("deleteMallGoodsInfo", mallGoodsInfoApi.DeleteGoodsInfo) // 刪除MallGoodsInfo
		mallGoodsInfoRouter.PUT("goods/status/:status", mallGoodsInfoApi.ChangeGoodsInfo)   // 上下架
		mallGoodsInfoRouter.PUT("goods", mallGoodsInfoApi.UpdateGoodsInfo)                  // 更新MallGoodsInfo
		mallGoodsInfoRouter.GET("goods/:id", mallGoodsInfoApi.FindGoodsInfo)                // 根據ID獲取MallGoodsInfo
		// mallGoodsInfoRouter.GET("goods/list", mallGoodsInfoApi.GetGoodsInfoList)               // 獲取MallGoodsInfo列表
	}
}
