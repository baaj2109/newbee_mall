package manage

import (
	v1 "github.com/baaj2109/newbee_mall/api/v1"
	"github.com/gin-gonic/gin"
)

type ManageIndexConfigRouter struct {
}

func (r *ManageIndexConfigRouter) InitManageIndexConfigRouter(Router *gin.RouterGroup) {
	mallIndexConfigRouter := Router.Group("v1")
	var mallIndexConfigApi = v1.ApiGroupApp.ManageApiGroup.ManageIndexConfigApi
	{
		mallIndexConfigRouter.POST("indexConfigs", mallIndexConfigApi.CreateIndexConfig)        // 新建MallIndexConfig
		mallIndexConfigRouter.POST("indexConfigs/delete", mallIndexConfigApi.DeleteIndexConfig) // 刪除MallIndexConfig
		mallIndexConfigRouter.PUT("indexConfigs", mallIndexConfigApi.UpdateIndexConfig)         // 更新MallIndexConfig
		mallIndexConfigRouter.GET("indexConfigs/:id", mallIndexConfigApi.FindIndexConfig)       // 根據ID獲取MallIndexConfig
		// mallIndexConfigRouter.GET("indexConfigs", mallIndexConfigApi.GetIndexConfigList)        // 獲取MallIndexConfig列表
	}
}
