package mall

import (
	"strconv"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model/request/order_request"
	"github.com/baaj2109/newbee_mall/model/response"
	"github.com/baaj2109/newbee_mall/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MallOrderApi struct{}

// / save order using gin context
func (*MallOrderApi) SaveOrder(c *gin.Context) {
	var saveOrderParam order_request.SaveOrderParam
	_ = c.ShouldBindJSON(&saveOrderParam)

	if err := utils.Verify(saveOrderParam, utils.SaveOrderParamVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	token := c.GetHeader("token")
	priceTotal := 0
	itemsForSave, _ := mallShopCartService.GetCartItemsForSettle(token, saveOrderParam.CartItemIds)
	if len(itemsForSave) < 1 {
		response.FailWithMessage("no items for save", c)
	} else {
		/// total price
		for _, shoppingCartItemVo := range itemsForSave {
			priceTotal += priceTotal + shoppingCartItemVo.GoodsCount*shoppingCartItemVo.SellingPrice
		}
		if priceTotal < 1 {
			response.FailWithMessage("invalid price", c)
		}
		userAddress, _ := mallUserAddressService.GetUserDefaultAddress(token)
		saveOrderResult, err := mallOrderService.SaveOrder(token, userAddress, itemsForSave)
		if err != nil {
			global.GVA_LOG.Error("failed to generate order", zap.Error(err))
			response.FailWithMessage("failed to generate order", c)
		}
		response.OkWithData(saveOrderResult, c)
	}
}

// / pay success using gin context
func (*MallOrderApi) PaySuccess(c *gin.Context) {
	orderNo := c.Query("orderNo")
	payType, _ := strconv.Atoi(c.Query("payType"))
	if err := mallOrderService.PaySuccess(orderNo, payType); err != nil {
		global.GVA_LOG.Error("failed to pay success", zap.Error(err))
		response.FailWithMessage("failed to pay success", c)
	}
	response.OkWithMessage("pay success", c)
}

// / finish order using gin context
func (*MallOrderApi) FinishOrder(c *gin.Context) {
	orderNo := c.Param("orderNo")
	token := c.GetHeader("token")
	if err := mallOrderService.FinishOrder(token, orderNo); err != nil {
		global.GVA_LOG.Error("failed to finish order", zap.Error(err))
		response.FailWithMessage("failed to finish order", c)
	}
	response.OkWithMessage("finish order success", c)
}

// / cancel order using gin context
func (*MallOrderApi) CancelOrder(c *gin.Context) {
	orderNo := c.Param("orderNo")
	token := c.GetHeader("token")
	if err := mallOrderService.CancelOrder(token, orderNo); err != nil {
		global.GVA_LOG.Error("failed to cancel order", zap.Error(err))
		response.FailWithMessage("failed to cancel order", c)
	}
	response.OkWithMessage("cancel order success", c)
}

// / get order detail using gin context
func (*MallOrderApi) OrderDetailPage(c *gin.Context) {
	orderNo := c.Param("orderNo")
	token := c.GetHeader("token")
	orderDetail, err := mallOrderService.GetOrderDetailByOrderNo(token, orderNo)
	if err != nil {
		global.GVA_LOG.Error("failed to get order detail", zap.Error(err))
		response.FailWithMessage("failed to get order detail", c)
	}
	response.OkWithData(orderDetail, c)
}

// / get order list using gin context
func (*MallOrderApi) OrderListPage(c *gin.Context) {
	token := c.GetHeader("token")
	pageNumber, _ := strconv.Atoi(c.Query("pageNumber"))
	status := c.Query("status")
	if pageNumber <= 0 {
		pageNumber = 1
	}
	if list, total, err := mallOrderService.MallOrderListBySearch(token, pageNumber, status); err != nil {
		global.GVA_LOG.Error("failed to get order list", zap.Error(err))
		response.FailWithMessage("failed to get order list", c)
	} else if len(list) < 1 {
		response.OkWithDetailed(response.PageResult{
			List:       make([]interface{}, 0),
			TotalCount: total,
			CurrPage:   pageNumber,
			PageSize:   5,
		}, "SUCCESS", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageNumber,
			PageSize:   5,
		}, "SUCCESS", c)
	}
}
