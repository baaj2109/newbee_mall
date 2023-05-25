package mall

import (
	"errors"
	"time"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model"
	"github.com/baaj2109/newbee_mall/model/common/enum"
	"github.com/baaj2109/newbee_mall/model/request/goods_info_request"
	"github.com/baaj2109/newbee_mall/model/response"
	"github.com/baaj2109/newbee_mall/utils"
	"github.com/jinzhu/copier"
)

type MallOrderService struct {
}

// save order using token, user address, cart item response
func (m *MallOrderService) SaveOrder(token string, userAddress model.UserAddress, myShoppingCartItems []response.CartItemResponse) (orderNo string, err error) {
	var userToken model.UserToken

	/// check user token
	err = global.GVA_DB.Where("token = ?", token).First(&userToken).Error
	if err != nil {
		return orderNo, err
	}
	var itemIdList []int
	var goodsIds []int
	for _, cartItem := range myShoppingCartItems {
		itemIdList = append(itemIdList, cartItem.CartItemId)
		goodsIds = append(goodsIds, cartItem.GoodsId)
	}
	var newMallGoods []model.GoodsInfo
	global.GVA_DB.Where("goods_id in ? ", goodsIds).Find(&newMallGoods)
	/// check is there any sold out goods in cart
	for _, mallGoods := range newMallGoods {
		if mallGoods.GoodsSellStatus != enum.GOODS_UNDER.Code() {
			return orderNo, errors.New("this item is sold out")
		}
	}

	newMallGoodsMap := make(map[int]model.GoodsInfo)
	for _, mallGoods := range newMallGoods {
		newMallGoodsMap[mallGoods.GoodsId] = mallGoods
	}

	for _, shoppingCartItemVO := range myShoppingCartItems {
		// return error if goods not inside cart
		if _, ok := newMallGoodsMap[shoppingCartItemVO.GoodsId]; !ok {
			return orderNo, errors.New("cart data invalid")
		}
		if shoppingCartItemVO.GoodsCount > newMallGoodsMap[shoppingCartItemVO.GoodsId].StockNum {
			return orderNo, errors.New("item is sold out")
		}
	}

	// delete
	if len(itemIdList) <= 0 || len(goodsIds) <= 0 {
		return
	}

	if err = global.GVA_DB.Where("cart_item_id in ?", itemIdList).Updates(
		model.ShoppingCartItem{
			IsDeleted: 1,
		}).Error; err != nil {
		return
	}

	var stockNumDTOS []goods_info_request.StockNumDTO
	copier.Copy(&stockNumDTOS, &myShoppingCartItems)
	for _, stockNumDTO := range stockNumDTOS {
		var goodsInfo model.GoodsInfo
		global.GVA_DB.Where("goods_id =?", stockNumDTO.GoodsId).First(&goodsInfo)
		if err = global.GVA_DB.Where("goods_id =? and stock_num>= ? and goods_sell_status = 0",
			stockNumDTO.GoodsId, stockNumDTO.GoodsCount).Updates(
			model.GoodsInfo{
				StockNum: goodsInfo.StockNum - stockNumDTO.GoodsCount,
			}).Error; err != nil {
			return orderNo, errors.New("item is sold out")
		}
	}

	/// generate order no
	orderNo = utils.GenOrderNo()
	priceTotal := 0

	/// save order
	var newBeeMallOrder model.Order
	newBeeMallOrder.OrderNo = orderNo
	newBeeMallOrder.UserId = userToken.UserId

	/// measure price
	for _, newBeeMallShoppingCartItemVO := range myShoppingCartItems {
		priceTotal = priceTotal + newBeeMallShoppingCartItemVO.GoodsCount*newBeeMallShoppingCartItemVO.SellingPrice
	}
	if priceTotal < 1 {
		return orderNo, errors.New("price invalid")
	}

	newBeeMallOrder.CreateTime = utils.JSONTime{Time: time.Now()}
	newBeeMallOrder.UpdateTime = utils.JSONTime{Time: time.Now()}
	newBeeMallOrder.TotalPrice = priceTotal
	newBeeMallOrder.ExtraInfo = ""

	/// generate order and save
	if err = global.GVA_DB.Save(&newBeeMallOrder).Error; err != nil {
		return orderNo, errors.New("failed to save orde to dataabase")
	}

	/// generate order address, and save it to database
	var newBeeMallOrderAddress model.OrderAddress
	copier.Copy(&newBeeMallOrderAddress, &userAddress)
	newBeeMallOrderAddress.OrderId = newBeeMallOrder.OrderId

	// generate order items, and save it to database
	var newBeeMallOrderItems []model.OrderItem
	for _, newBeeMallShoppingCartItemVO := range myShoppingCartItems {
		var newBeeMallOrderItem model.OrderItem
		copier.Copy(&newBeeMallOrderItem, &newBeeMallShoppingCartItemVO)
		newBeeMallOrderItem.OrderId = newBeeMallOrder.OrderId
		newBeeMallOrderItem.CreateTime = utils.JSONTime{Time: time.Now()}
		newBeeMallOrderItems = append(newBeeMallOrderItems, newBeeMallOrderItem)
	}
	if err = global.GVA_DB.Save(&newBeeMallOrderItems).Error; err != nil {
		return orderNo, err
	}
	return orderNo, nil
}
