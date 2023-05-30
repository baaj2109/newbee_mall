package model

import "github.com/baaj2109/newbee_mall/utils"

type ShoppingCartItem struct {
	CartItemId int            `json:"cartItemId" form:"cartItemId" gorm:"primarykey;AUTO_INCREMENT"`
	UserId     int            `json:"userId" form:"userId" gorm:"column:user_id;comment:用戶主键id;type:bigint"`
	GoodsId    int            `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:關聯商品id;type:bigint"`
	GoodsCount int            `json:"goodsCount" form:"goodsCount" gorm:"column:goods_count;comment:數量(最大為5);type:int"`
	IsDeleted  int            `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:刪除標識字段(0-未刪除 1-已刪除);type:tinyint"`
	CreateTime utils.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:創建時間;type:datetime"`
	UpdateTime utils.JSONTime `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:最新修改時間;type:datetime"`
}

// TableName MallShoppingCartItem 表名
func (ShoppingCartItem) TableName() string {
	return "tb_newbee_mall_shopping_cart_item"
}
