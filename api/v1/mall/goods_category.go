package mall

import (
	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MallGoodsCategoryApi struct {
}

// 返回分類數據(首頁調用)
func (m *MallGoodsCategoryApi) GetGoodsCategory(c *gin.Context) {
	list, err := mallGoodsCategoryService.GetCategoriesForIndex()
	if err != nil {
		global.GVA_LOG.Error("invalid search!", zap.Error(err))
		response.FailWithMessage("invalid search "+err.Error(), c)
	}
	response.OkWithData(list, c)
}
