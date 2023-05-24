package manage

import (
	"strconv"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model"
	"github.com/baaj2109/newbee_mall/model/request"
	"github.com/baaj2109/newbee_mall/model/request/goods_info_request"
	"github.com/baaj2109/newbee_mall/model/request/update_request"
	"github.com/baaj2109/newbee_mall/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ManageGoodsInfoApi struct {
}

func (m *ManageGoodsInfoApi) CreateGoodsInfo(c *gin.Context) {
	var param goods_info_request.GoodsInfoAddParam
	_ = c.ShouldBindJSON(&param)
	if err := mallGoodsInfoService.CreateGoodsInfo(param); err != nil {
		global.GVA_LOG.Error("create goods info failed:", zap.Error(err))
		response.FailWithMessage("create goods info failed:"+err.Error(), c)
	}
	response.OkWithMessage("create goods info success", c)
}

func (m *ManageGoodsInfoApi) UpdateGoodsInfo(c *gin.Context) {
	var param update_request.GoodsInfoUpdateParam
	_ = c.ShouldBindJSON(&param)
	if err := mallGoodsInfoService.UpdateGoodsInfo(param); err != nil {
		global.GVA_LOG.Error("update goods info failed:", zap.Error(err))
		response.FailWithMessage("update goods info failed:"+err.Error(), c)
	}
	response.OkWithMessage("update goods info success", c)
}

func (m *ManageGoodsInfoApi) DeleteGoodsInfo(c *gin.Context) {
	var goodsInfo model.GoodsInfo
	_ = c.ShouldBindJSON(&goodsInfo)
	if err := mallGoodsInfoService.DeleteGoodsInfo(goodsInfo); err != nil {
		global.GVA_LOG.Error("delete goods info failed:", zap.Error(err))
		response.FailWithMessage("delete goods info failed:"+err.Error(), c)
	}
	response.OkWithMessage("delete goods info success", c)
}

// change goods info by ids
func (m *ManageGoodsInfoApi) ChangeGoodsInfo(c *gin.Context) {
	var ids request.IdsReq
	_ = c.ShouldBindJSON(&ids)
	sellStatus := c.Param("status")
	if err := mallGoodsInfoService.ChangeMallGoodsInfoByIds(ids, sellStatus); err != nil {
		global.GVA_LOG.Error("change goods info failed:", zap.Error(err))
		response.FailWithMessage("change goods info failed:"+err.Error(), c)
	}
	response.OkWithMessage("change goods info success", c)
}

// find goods info use gin.context param id
func (m *ManageGoodsInfoApi) FindGoodsInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	goodsInfo, err := mallGoodsInfoService.GetGoodsInfo(id)
	if err != nil {
		global.GVA_LOG.Error("find goods info failed:", zap.Error(err))
		response.FailWithMessage("find goods info failed:"+err.Error(), c)
	}
	goodsInfoRes := make(map[string]interface{})
	goodsInfoRes["goods"] = goodsInfo
	if thirdCategory, _ := mallGoodsCategoryService.SelectCategoryById(goodsInfo.GoodsCategoryId); thirdCategory != (model.GoodsCategory{}) {
		goodsInfoRes["thirdCategory"] = thirdCategory
		if secondCategory, _ := mallGoodsCategoryService.SelectCategoryById(thirdCategory.ParentId); secondCategory != (model.GoodsCategory{}) {
			goodsInfoRes["secondCategory"] = secondCategory
			if firstCategory, _ := mallGoodsCategoryService.SelectCategoryById(secondCategory.ParentId); firstCategory != (model.GoodsCategory{}) {
				goodsInfoRes["firstCategory"] = firstCategory
			}
		}
	}
	response.OkWithData(goodsInfoRes, c)
}
