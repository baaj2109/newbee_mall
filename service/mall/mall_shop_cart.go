package mall

import (
	"errors"
	"time"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model"
	"github.com/baaj2109/newbee_mall/model/request/shop_cart_request"
	"github.com/baaj2109/newbee_mall/model/request/update_request"
	"github.com/baaj2109/newbee_mall/model/response"
	"github.com/baaj2109/newbee_mall/utils"
	"github.com/jinzhu/copier"
)

type MallShopCartService struct {
}

func (m *MallShopCartService) GetShoppingCartItems(token string) (cartItems []response.CartItemResponse, err error) {
	var userToken model.UserToken
	if err = global.GVA_DB.Where("token = ?", token).First(&userToken).Error; err != nil {
		return cartItems, errors.New("token invalid")
	}
	var shopCartItems []model.ShoppingCartItem
	var goodsInfos []model.GoodsInfo
	global.GVA_DB.Where("user_id=? and is_deleted = 0", userToken.UserId).Find(&shopCartItems)
	var goodsIds []int
	for _, shopCartItem := range shopCartItems {
		goodsIds = append(goodsIds, shopCartItem.GoodsId)
	}
	global.GVA_DB.Where("goods_id in ?", goodsIds).Find(&goodsInfos)
	goodsMap := make(map[int]model.GoodsInfo)
	for _, goodsInfo := range goodsInfos {
		goodsMap[goodsInfo.GoodsId] = goodsInfo
	}
	for _, v := range shopCartItems {
		var cartItem response.CartItemResponse
		copier.Copy(&cartItem, &v)
		if _, ok := goodsMap[v.GoodsId]; ok {
			goodsInfo := goodsMap[v.GoodsId]
			cartItem.GoodsName = goodsInfo.GoodsName
			cartItem.GoodsCoverImg = goodsInfo.GoodsCoverImg
			cartItem.SellingPrice = goodsInfo.SellingPrice
		}
		cartItems = append(cartItems, cartItem)
	}
	return
}
func (m *MallShopCartService) SaveMallCartItem(token string, req shop_cart_request.SaveCartItemParam) (err error) {
	if req.GoodsCount < 1 {
		return errors.New("invalid goods count")
	}
	if req.GoodsCount > 5 {
		return errors.New("invalid goods count")
	}
	var userToken model.UserToken
	if err = global.GVA_DB.Where("token = ?", token).First(&userToken).Error; err != nil {
		return errors.New("token invalid")
	}
	var shopCartItems []model.ShoppingCartItem
	if err = global.GVA_DB.Where("user_id = ? and goods_id = ? and is_deleted = 0",
		userToken.UserId, req.GoodsId).Find(&shopCartItems).Error; err != nil {
		return errors.New("goods already exists")
	}
	if err = global.GVA_DB.Where("goods_id = ? ",
		req.GoodsId).First(&model.GoodsInfo{}).Error; err != nil {
		return errors.New("goods is empty")
	}
	var total int64
	global.GVA_DB.Where("user_id =?  and is_deleted = 0", userToken.UserId).Count(&total)
	if total > 20 {
		return errors.New("invalid goods count")
	}
	var shopCartItem model.ShoppingCartItem
	if err = copier.Copy(&shopCartItem, &req); err != nil {
		return err
	}
	shopCartItem.UserId = userToken.UserId
	shopCartItem.CreateTime = utils.JSONTime{Time: time.Now()}
	shopCartItem.UpdateTime = utils.JSONTime{Time: time.Now()}
	err = global.GVA_DB.Save(&shopCartItem).Error
	return
}

func (m *MallShopCartService) UpdateMallCartItem(token string, req update_request.UpdateCartItemParam) (err error) {
	if req.GoodsCount > 5 {
		return errors.New("exceed max goods count")
	}
	var userToken model.UserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("token invalid")
	}
	var shopCartItem model.ShoppingCartItem
	if err = global.GVA_DB.Where("cart_item_id=? and is_deleted = 0", req.CartItemId).First(&shopCartItem).Error; err != nil {
		return errors.New("cart item not exists")
	}
	if shopCartItem.UserId != userToken.UserId {
		return errors.New("token invalid")
	}
	shopCartItem.GoodsCount = req.GoodsCount
	shopCartItem.UpdateTime = utils.JSONTime{Time: time.Now()}
	err = global.GVA_DB.Save(&shopCartItem).Error
	return
}

func (m *MallShopCartService) DeleteMallCartItem(token string, id int) (err error) {
	var userToken model.UserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("token invalid")
	}
	var shopCartItem model.ShoppingCartItem
	err = global.GVA_DB.Where("cart_item_id = ? and is_deleted = 0", id).First(&shopCartItem).Error
	if err != nil {
		return
	}
	if userToken.UserId != shopCartItem.UserId {
		return errors.New("token invalid")
	}
	err = global.GVA_DB.Where("cart_item_id = ? and is_deleted = 0", id).UpdateColumns(
		&model.ShoppingCartItem{
			IsDeleted: 1,
		}).Error
	return
}

func (m *MallShopCartService) GetCartItemsForSettle(token string, cartItemIds []int) (cartItemRes []response.CartItemResponse, err error) {
	var userToken model.UserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return cartItemRes, errors.New("token invalid")
	}
	var shopCartItems []model.ShoppingCartItem
	err = global.GVA_DB.Where("cart_item_id in (?) and user_id = ? and is_deleted = 0", cartItemIds, userToken.UserId).Find(&shopCartItems).Error
	if err != nil {
		return
	}
	cartItemRes, _ = getMallShoppingCartItemVOS(shopCartItems)
	//measure total price
	priceTotal := 0
	for _, cartItem := range cartItemRes {
		priceTotal = priceTotal + cartItem.GoodsCount*cartItem.SellingPrice
	}
	return
}

func getMallShoppingCartItemVOS(cartItems []model.ShoppingCartItem) (cartItemsRes []response.CartItemResponse, err error) {
	var goodsIds []int
	for _, cartItem := range cartItems {
		goodsIds = append(goodsIds, cartItem.GoodsId)
	}
	var newBeeMallGoods []model.GoodsInfo
	err = global.GVA_DB.Where("goods_id in ?", goodsIds).Find(&newBeeMallGoods).Error
	if err != nil {
		return
	}

	newBeeMallGoodsMap := make(map[int]model.GoodsInfo)
	for _, goodsInfo := range newBeeMallGoods {
		newBeeMallGoodsMap[goodsInfo.GoodsId] = goodsInfo
	}
	for _, cartItem := range cartItems {
		var cartItemRes response.CartItemResponse
		copier.Copy(&cartItemRes, &cartItem)
		// is contain key
		if _, ok := newBeeMallGoodsMap[cartItemRes.GoodsId]; ok {
			newBeeMallGoodsTemp := newBeeMallGoodsMap[cartItemRes.GoodsId]
			cartItemRes.GoodsCoverImg = newBeeMallGoodsTemp.GoodsCoverImg
			goodsName := utils.SubStrLen(newBeeMallGoodsTemp.GoodsName, 28)
			cartItemRes.GoodsName = goodsName
			cartItemRes.SellingPrice = newBeeMallGoodsTemp.SellingPrice
			cartItemsRes = append(cartItemsRes, cartItemRes)
		}
	}
	return
}
