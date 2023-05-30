package manage

import (
	v1 "github.com/baaj2109/newbee_mall/api/v1"
	"github.com/baaj2109/newbee_mall/middleware"
	"github.com/gin-gonic/gin"
)

type ManageCarouselRouter struct {
}

func (r *ManageCarouselRouter) InitManageCarouselRouter(Router *gin.RouterGroup) {
	mallCarouselRouter := Router.Group("v1").Use(middleware.AdminJWTAuth())
	var mallCarouselApi = v1.ApiGroupApp.ManageApiGroup.ManageCarouselApi
	{
		mallCarouselRouter.POST("carousels", mallCarouselApi.CreateCarousel)   // 新建MallCarousel
		mallCarouselRouter.DELETE("carousels", mallCarouselApi.DeleteCarousel) // 刪除MallCarousel
		mallCarouselRouter.PUT("carousels", mallCarouselApi.UpdateCarousel)    // 更新MallCarousel
		mallCarouselRouter.GET("carousels/:id", mallCarouselApi.FindCarousel)  // 根據ID獲取輪播圖
		// mallCarouselRouter.GET("carousels", mallCarouselApi.GetCarouselList)   // 獲取輪播圖列表
	}
}
