package mall

import (
	"strconv"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MallGoodsInfoApi struct {
}

// / goods search
func (m *MallGoodsInfoApi) GoodsSearch(c *gin.Context) {
	pageNumber, _ := strconv.Atoi(c.Query("pageNumber"))
	goodsCategoryId, _ := strconv.Atoi(c.Query("goodsCategoryId"))
	keyword := c.Query("keyword")
	orderBy := c.Query("orderBy")
	if list, total, err := mallGoodsInfoService.MallGoodsListBySearch(pageNumber, goodsCategoryId, keyword, orderBy); err != nil {
		global.GVA_LOG.Error("invalid search", zap.Error(err))
		response.FailWithMessage("invalid search"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageNumber,
			PageSize:   10,
		}, "search success", c)
	}
}

func (m *MallGoodsInfoApi) GoodsDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	goodsInfo, err := mallGoodsInfoService.GetMallGoodsInfo(id)
	if err != nil {
		global.GVA_LOG.Error("get mall goods info failed!", zap.Error(err))
		response.FailWithMessage("get mall goods info failed"+err.Error(), c)
	}
	response.OkWithData(goodsInfo, c)
}
