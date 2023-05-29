package mall

import (
	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model/common/enum"
	"github.com/baaj2109/newbee_mall/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MallIndexApi struct {
}

func (api *MallIndexApi) MallIndex(c *gin.Context) {
	carouseInfo, _, err := mallCarouselService.GetCarouselsForIndex(5)
	if err != nil {
		global.GVA_LOG.Error("failed to get carousel info", zap.Error(err))
		response.FailWithMessage("failed to get carousel info", c)
	}
	hotGoodses, err := mallIndexConfigService.GetConfigGoodsForIndex(enum.IndexGoodsHot.Code(), 4)
	if err != nil {
		global.GVA_LOG.Error("failed to get hot goods info", zap.Error(err))
		response.FailWithMessage("failed to get hot goods info", c)
	}
	newGoodes, err := mallIndexConfigService.GetConfigGoodsForIndex(enum.IndexGoodsNew.Code(), 4)
	if err != nil {
		global.GVA_LOG.Error("failed to get new goods info", zap.Error(err))
		response.FailWithMessage("failed to get new goods info", c)
	}
	recommendGoodses, err := mallIndexConfigService.GetConfigGoodsForIndex(enum.IndexGoodsRecommond.Code(), 4)
	if err != nil {
		global.GVA_LOG.Error("failed to get recommend goods info", zap.Error(err))
		response.FailWithMessage("failed to get recommend goods info", c)
	}
	indexResult := make(map[string]interface{})
	indexResult["carousels"] = carouseInfo
	indexResult["hotGoodses"] = hotGoodses
	indexResult["newGoodses"] = newGoodes
	indexResult["recommendGoodses"] = recommendGoodses
	response.OkWithData(indexResult, c)

}
