package shop_cart_request

type SaveCartItemParam struct {
	GoodsCount int `json:"goodsCount"`
	GoodsId    int `json:"goodsId"`
}
