package model

import "github.com/baaj2109/newbee_mall/utils"

type OrderItem struct {
	OrderItemId   int            `json:"orderItemId" gorm:"primarykey;AUTO_INCREMENT"`
	OrderId       int            `json:"orderId" form:"orderId" gorm:"column:order_id;;type:bigint"`
	GoodsId       int            `json:"goodsId" form:"goodsId" gorm:"column:goods_id;;type:bigint"`
	GoodsName     string         `json:"goodsName" form:"goodsName" gorm:"column:goods_name;comment:商品名;type:varchar(200);"`
	GoodsCoverImg string         `json:"goodsCoverImg" form:"goodsCoverImg" gorm:"column:goods_cover_img;comment:商品主圖;type:varchar(200);"`
	SellingPrice  int            `json:"sellingPrice" form:"sellingPrice" gorm:"column:selling_price;comment:商品實際售價;type:int"`
	GoodsCount    int            `json:"goodsCount" form:"goodsCount" gorm:"column:goods_count;;type:bigint"`
	CreateTime    utils.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:創建時間;type:datetime"`
}

func (OrderItem) TableName() string {
	return "tb_newbee_mall_order_item"
}
