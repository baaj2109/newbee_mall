package response

import "github.com/baaj2109/newbee_mall/utils"

type MallOrderResponse struct {
	OrderId                int                     `json:"orderId"`
	OrderNo                string                  `json:"orderNo"`
	TotalPrice             int                     `json:"totalPrice"`
	PayType                int                     `json:"payType"`
	OrderStatus            int                     `json:"orderStatus"`
	OrderStatusString      string                  `json:"orderStatusString"`
	CreateTime             utils.JSONTime          `json:"createTime"`
	NewBeeMallOrderItemVOS []NewBeeMallOrderItemVO `json:"newBeeMallOrderItemVOS"`
}

type NewBeeMallOrderItemVO struct {
	GoodsId       int    `json:"goodsId"`
	GoodsName     string `json:"goodsName"`
	GoodsCount    int    `json:"goodsCount"`
	GoodsCoverImg string `json:"goodsCoverImg"`
	SellingPrice  int    `json:"sellingPrice"`
}

type MallOrderDetailVO struct {
	OrderNo                string                  `json:"orderNo"`
	TotalPrice             int                     `json:"totalPrice"`
	PayStatus              int                     `json:"payStatus"`
	PayType                int                     `json:"payType"`
	PayTypeString          string                  `json:"payTypeString"`
	PayTime                utils.JSONTime          `json:"payTime"`
	OrderStatus            int                     `json:"orderStatus"`
	OrderStatusString      string                  `json:"orderStatusString"`
	CreateTime             utils.JSONTime          `json:"createTime"`
	NewBeeMallOrderItemVOS []NewBeeMallOrderItemVO `json:"newBeeMallOrderItemVOS"`
}
