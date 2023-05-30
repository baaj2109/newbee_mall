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

func (m *MallOrderService) PaySuccess(orderNo string, payType int) (err error) {
	var mallOrder model.Order
	err = global.GVA_DB.Where("order_no = ? and is_deleted=0 ", orderNo).First(&mallOrder).Error
	if mallOrder != (model.Order{}) {
		if mallOrder.OrderStatus != 0 {
			return errors.New("invalid order status")
		}
		mallOrder.OrderStatus = enum.ORDER_PAID.Code()
		mallOrder.PayType = payType
		mallOrder.PayStatus = 1
		mallOrder.PayTime = utils.JSONTime{Time: time.Now()}
		mallOrder.UpdateTime = utils.JSONTime{Time: time.Now()}
		err = global.GVA_DB.Save(&mallOrder).Error
	}
	return
}

func (m *MallOrderService) FinishOrder(token string, orderNo string) (err error) {
	var userToken model.UserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("user can't find")
	}
	var mallOrder model.Order
	if err = global.GVA_DB.Where("order_no=? and is_deleted = 0", orderNo).First(&mallOrder).Error; err != nil {
		return errors.New("order can't find")
	}
	if mallOrder.UserId != userToken.UserId {
		return errors.New("user can't find")
	}
	mallOrder.OrderStatus = enum.ORDER_SUCCESS.Code()
	mallOrder.UpdateTime = utils.JSONTime{Time: time.Now()}
	err = global.GVA_DB.Save(&mallOrder).Error
	return
}
func (m *MallOrderService) CancelOrder(token string, orderNo string) (err error) {
	var userToken model.UserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("user can't find")
	}
	var mallOrder model.Order
	if err = global.GVA_DB.Where("order_no=? and is_deleted = 0", orderNo).First(&mallOrder).Error; err != nil {
		return errors.New("order can't find")
	}
	if mallOrder.UserId != userToken.UserId {
		return errors.New("invalid user")
	}
	if utils.NumsInList(mallOrder.OrderStatus, []int{enum.ORDER_SUCCESS.Code(),
		enum.ORDER_CLOSED_BY_MALLUSER.Code(), enum.ORDER_CLOSED_BY_EXPIRED.Code(), enum.ORDER_CLOSED_BY_JUDGE.Code()}) {
		return errors.New("invalid order status")
	}
	mallOrder.OrderStatus = enum.ORDER_CLOSED_BY_MALLUSER.Code()
	mallOrder.UpdateTime = utils.JSONTime{Time: time.Now()}
	err = global.GVA_DB.Save(&mallOrder).Error
	return
}

func (m *MallOrderService) GetOrderDetailByOrderNo(token string, orderNo string) (orderDetail response.MallOrderDetailVO, err error) {
	var userToken model.UserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return orderDetail, errors.New("user can't find")
	}
	var mallOrder model.Order
	if err = global.GVA_DB.Where("order_no=? and is_deleted = 0", orderNo).First(&mallOrder).Error; err != nil {
		return orderDetail, errors.New("order can't find")
	}
	if mallOrder.UserId != userToken.UserId {
		return orderDetail, errors.New("invalid user")
	}
	var orderItems []model.OrderItem
	err = global.GVA_DB.Where("order_id = ?", mallOrder.OrderId).Find(&orderItems).Error
	if len(orderItems) <= 0 {
		return orderDetail, errors.New("item can't find")
	}

	var newBeeMallOrderItemVOS []response.NewBeeMallOrderItemVO
	copier.Copy(&newBeeMallOrderItemVOS, &orderItems)
	copier.Copy(&orderDetail, &mallOrder)
	// 訂單狀態前端顯示為中文
	_, OrderStatusStr := enum.GetNewBeeMallOrderStatusEnumByStatus(orderDetail.OrderStatus)
	_, payTapStr := enum.GetNewBeeMallOrderStatusEnumByStatus(orderDetail.PayType)
	orderDetail.OrderStatusString = OrderStatusStr
	orderDetail.PayTypeString = payTapStr
	orderDetail.NewBeeMallOrderItemVOS = newBeeMallOrderItemVOS

	return
}

func (m *MallOrderService) MallOrderListBySearch(token string, pageNumber int, status string) (list []response.MallOrderResponse, total int64, err error) {
	var userToken model.UserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return list, total, errors.New("user can't find")
	}
	// 根據搜索条件查詢
	var newBeeMallOrders []model.Order
	db := global.GVA_DB.Model(&newBeeMallOrders)

	if status != "" {
		db.Where("order_status = ?", status)
	}
	err = db.Where("user_id =? and is_deleted=0 ", userToken.UserId).Count(&total).Error
	//这里前段没有做滚動加载，直接顯示全部訂單
	//limit := 5
	offset := 5 * (pageNumber - 1)
	err = db.Offset(offset).Order(" order_id desc").Find(&newBeeMallOrders).Error

	var orderListVOS []response.MallOrderResponse
	if total > 0 {
		//數據轉换 将实体類轉成vo
		copier.Copy(&orderListVOS, &newBeeMallOrders)
		//设置訂單狀態中文顯示值
		for _, newBeeMallOrderListVO := range orderListVOS {
			_, statusStr := enum.GetNewBeeMallOrderStatusEnumByStatus(newBeeMallOrderListVO.OrderStatus)
			newBeeMallOrderListVO.OrderStatusString = statusStr
		}
		// 返回訂單id
		var orderIds []int
		for _, order := range newBeeMallOrders {
			orderIds = append(orderIds, order.OrderId)
		}
		//獲取OrderItem
		var orderItems []model.OrderItem
		if len(orderIds) > 0 {
			global.GVA_DB.Where("order_id in ?", orderIds).Find(&orderItems)
			itemByOrderIdMap := make(map[int][]model.OrderItem)
			for _, orderItem := range orderItems {
				itemByOrderIdMap[orderItem.OrderId] = []model.OrderItem{}
			}
			for k, v := range itemByOrderIdMap {
				for _, orderItem := range orderItems {
					if k == orderItem.OrderId {
						v = append(v, orderItem)
					}
					itemByOrderIdMap[k] = v
				}
			}
			//封装每個訂單列表对象的訂單項數據
			for _, newBeeMallOrderListVO := range orderListVOS {
				if _, ok := itemByOrderIdMap[newBeeMallOrderListVO.OrderId]; ok {
					orderItemListTemp := itemByOrderIdMap[newBeeMallOrderListVO.OrderId]
					var newBeeMallOrderItemVOS []response.NewBeeMallOrderItemVO
					copier.Copy(&newBeeMallOrderItemVOS, &orderItemListTemp)
					newBeeMallOrderListVO.NewBeeMallOrderItemVOS = newBeeMallOrderItemVOS
					_, OrderStatusStr := enum.GetNewBeeMallOrderStatusEnumByStatus(newBeeMallOrderListVO.OrderStatus)
					newBeeMallOrderListVO.OrderStatusString = OrderStatusStr
					list = append(list, newBeeMallOrderListVO)
				}
			}
		}
	}
	return list, total, err
}
