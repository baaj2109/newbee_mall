package manage

import (
	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model/common/request"
	"github.com/baaj2109/newbee_mall/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ManageOrderApi struct {
}

// check done order by gin context
func (m *ManageOrderApi) CheckDoneOrder(c *gin.Context) {
	var ids request.IdsReq
	_ = c.ShouldBindJSON(&ids)
	if err := mallOrderService.CheckDone(ids); err != nil {
		global.GVA_LOG.Error("failed to check done order", zap.Error(err))
		response.FailWithMessage("failed to check done order"+err.Error(), c)
	}
	response.OkWithMessage("success to check done order", c)
}

// check out order by gin context
func (m *ManageOrderApi) CheckOutOrder(c *gin.Context) {
	var ids request.IdsReq
	_ = c.ShouldBindJSON(&ids)
	if err := mallOrderService.CheckOut(ids); err != nil {
		global.GVA_LOG.Error("failed to check out order", zap.Error(err))
		response.FailWithMessage("failed to check out order"+err.Error(), c)
	}
	response.OkWithMessage("success to check out order", c)
}

// /close order by gin context
func (m *ManageOrderApi) CloseOrder(c *gin.Context) {
	var ids request.IdsReq
	_ = c.ShouldBindJSON(&ids)
	if err := mallOrderService.CloseOrder(ids); err != nil {
		global.GVA_LOG.Error("failed to close order", zap.Error(err))
		response.FailWithMessage("failed to close order"+err.Error(), c)
	}
	response.OkWithMessage("success to close order", c)
}
