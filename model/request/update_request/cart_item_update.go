package update_request

type UpdateCartItemParam struct {
	CartItemId int `json:"cartItemId"`
	GoodsCount int `json:"goodsCount"`
}
