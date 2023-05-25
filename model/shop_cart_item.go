package model

import "github.com/baaj2109/newbee_mall/utils"

type ShoppingCartItem struct {
	CartItemId int            `json:"cartItemId" form:"cartItemId" gorm:"primarykey;AUTO_INCREMENT"`
	UserId     int            `json:"userId" form:"userId" gorm:"column:user_id;comment:用户主键id;type:bigint"`
	GoodsId    int            `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:关联商品id;type:bigint"`
	GoodsCount int            `json:"goodsCount" form:"goodsCount" gorm:"column:goods_count;comment:数量(最大为5);type:int"`
	IsDeleted  int            `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:删除标识字段(0-未删除 1-已删除);type:tinyint"`
	CreateTime utils.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:datetime"`
	UpdateTime utils.JSONTime `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:最新修改时间;type:datetime"`
}

// TableName MallShoppingCartItem 表名
func (ShoppingCartItem) TableName() string {
	return "tb_newbee_mall_shopping_cart_item"
}
