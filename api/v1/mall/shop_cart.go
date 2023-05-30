package mall

import (
	"strconv"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model/request/shop_cart_request"
	"github.com/baaj2109/newbee_mall/model/request/update_request"
	"github.com/baaj2109/newbee_mall/model/response"
	"github.com/baaj2109/newbee_mall/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MallShopCartApi struct {
}

// / get cart item list using gin context
func (m *MallShopCartApi) GetCartItemList(c *gin.Context) {
	token := c.GetHeader("token")
	res, err := mallShopCartService.GetShoppingCartItems(token)
	if err != nil {
		global.GVA_LOG.Error("get cart item list failed", zap.Error(err))
		response.FailWithMessage("get cart item list failed", c)
	}
	response.OkWithData(res, c)
}

// / save cart item using gin context
func (m *MallShopCartApi) SaveCartItem(c *gin.Context) {
	var request shop_cart_request.SaveCartItemParam
	_ = c.ShouldBindJSON(&request)
	token := c.GetHeader("token")
	if err := mallShopCartService.SaveMallCartItem(token, request); err != nil {
		global.GVA_LOG.Error("save cart item failed", zap.Error(err))
		response.FailWithMessage("save cart item failed", c)
	} else {
		response.OkWithMessage("save cart item success", c)
	}
}

// / update cart item using gin context
func (m *MallShopCartApi) UpdateCartItem(c *gin.Context) {
	var request update_request.UpdateCartItemParam
	_ = c.ShouldBindJSON(&request)
	token := c.GetHeader("token")
	if err := mallShopCartService.UpdateMallCartItem(token, request); err != nil {
		global.GVA_LOG.Error("update cart item failed", zap.Error(err))
		response.FailWithMessage("update cart item failed", c)
	} else {
		response.OkWithMessage("update cart item success", c)
	}
}

// / delete cart item using gin context
func (m *MallShopCartApi) DeleteCartItem(c *gin.Context) {
	token := c.GetHeader("token")
	id, _ := strconv.Atoi(c.Param("newBeeMallShoppingCartItemId"))
	if err := mallShopCartService.DeleteMallCartItem(token, id); err != nil {
		global.GVA_LOG.Error("delete cart item failed", zap.Error(err))
		response.FailWithMessage("delete cart item failed", c)
	} else {
		response.OkWithMessage("delete cart item success", c)
	}
}

// / to settle
func (m *MallShopCartApi) ToSettle(c *gin.Context) {
	cartItemIds := c.Query("cartItemIds")
	token := c.GetHeader("token")
	ids := utils.StrToInt(cartItemIds)
	if res, err := mallShopCartService.GetCartItemsForSettle(token, ids); err != nil {
		global.GVA_LOG.Error("to settle failed", zap.Error(err))
		response.FailWithMessage("to settle failed"+err.Error(), c)
	} else {
		response.OkWithData(res, c)
	}
}
