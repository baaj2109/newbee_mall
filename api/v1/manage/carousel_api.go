package manage

import (
	"strconv"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model/request"
	"github.com/baaj2109/newbee_mall/model/request/carousel_request"
	"github.com/baaj2109/newbee_mall/model/request/update_request"
	"github.com/baaj2109/newbee_mall/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ManageCarouselApi struct {
}

func (m *ManageCarouselApi) CreateCarousel(c *gin.Context) {
	var param carousel_request.MallCarouselAddParam
	_ = c.ShouldBindJSON(&param)
	if err := mallCarouselService.CreateCarousel(param); err != nil {
		global.GVA_LOG.Error("create carousel failed:", zap.Error(err))
		response.FailWithMessage("create carousel failed:"+err.Error(), c)
	} else {
		response.OkWithMessage("create carousel success", c)
	}
}

func (m *ManageCarouselApi) DeleteCarousel(c *gin.Context) {
	var ids request.IdsReq
	_ = c.ShouldBindJSON(&ids)
	if err := mallCarouselService.DeleteCarousel(ids); err != nil {
		global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete"+err.Error(), c)
	} else {
		response.OkWithMessage("delete success", c)
	}
}

func (m *ManageCarouselApi) UpdateCarousel(c *gin.Context) {
	var param update_request.CarouselUpdateParam
	_ = c.ShouldBindJSON(&param)
	if err := mallCarouselService.UpdateCarousel(param); err != nil {
		global.GVA_LOG.Error("failed to update!", zap.Error(err))
		response.FailWithMessage("failed to update:"+err.Error(), c)
	} else {
		response.OkWithMessage("update success", c)
	}
}

func (m *ManageCarouselApi) FindCarousel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if mallCarousel, err := mallCarouselService.GetCarousel(id); err != nil {
		global.GVA_LOG.Error("fail to find carousel!"+err.Error(), zap.Error(err))
		response.FailWithMessage("failed to find carousel", c)
	} else {
		response.OkWithData(mallCarousel, c)
	}
}
