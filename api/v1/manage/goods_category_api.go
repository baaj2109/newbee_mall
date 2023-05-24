package manage

import (
	"strconv"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model/request"
	"github.com/baaj2109/newbee_mall/model/request/goods_category_request"
	"github.com/baaj2109/newbee_mall/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ManageGoodsCategoryApi struct {
}

// create goods category by gin context
func (m *ManageGoodsCategoryApi) CreateCategory(c *gin.Context) {
	var request goods_category_request.GoodsCategoryReq
	_ = c.ShouldBindJSON(&request)
	if err := mallGoodsCategoryService.CreateCategory(request); err != nil {
		global.GVA_LOG.Error("create goods category failed:", zap.Error(err))
		response.FailWithMessage("create goods category failed:"+err.Error(), c)
	} else {
		response.OkWithMessage("create goods category success", c)
	}
}

// update goods category by gin context
func (m *ManageGoodsCategoryApi) UpdateCategory(c *gin.Context) {
	var param goods_category_request.GoodsCategoryReq
	_ = c.ShouldBindJSON(&param)
	if err := mallGoodsCategoryService.UpdateCategory(param); err != nil {
		global.GVA_LOG.Error("update goods category failed:", zap.Error(err))
		response.FailWithMessage("update goods category failed, already got category", c)
	} else {
		response.OkWithMessage("update goods category success", c)
	}
}

// delete goods category by gin context
func (m *ManageGoodsCategoryApi) DelCategory(c *gin.Context) {
	var ids request.IdsReq
	_ = c.ShouldBindJSON(&ids)
	if _, err := mallGoodsCategoryService.DeleteGoodsCategoriesByIds(ids); err != nil {
		global.GVA_LOG.Error("delete goods category failed:", zap.Error(err))
		response.FailWithMessage("delete goods category failed:"+err.Error(), c)
	} else {
		response.OkWithMessage("delete goods category success", c)
	}
}

// get goods category by gin context
func (m *ManageGoodsCategoryApi) GetCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	goodsCategory, err := mallGoodsCategoryService.SelectCategoryById(id)
	if err != nil {
		global.GVA_LOG.Error("get goods category failed:", zap.Error(err))
		response.FailWithMessage("get goods category failed:"+err.Error(), c)
	} else {
		response.OkWithData(response.GoodsCategoryResponse{GoodsCategory: goodsCategory}, c)
	}
}
