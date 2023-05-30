package model

import "github.com/baaj2109/newbee_mall/utils"

type Order struct {
	OrderId     int            `json:"orderId" form:"orderId" gorm:"primarykey;AUTO_INCREMENT"`
	OrderNo     string         `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:訂單號;type:varchar(20);"`
	UserId      int            `json:"userId" form:"userId" gorm:"column:user_id;comment:用戶主键id;type:bigint"`
	TotalPrice  int            `json:"totalPrice" form:"totalPrice" gorm:"column:total_price;comment:訂單總價;type:int"`
	PayStatus   int            `json:"payStatus" form:"payStatus" gorm:"column:pay_status;comment:支付狀態:0.未支付,1.支付成功,-1:支付失敗;type:tinyint"`
	PayType     int            `json:"payType" form:"payType" gorm:"column:pay_type;comment:0.无 1.支付宝支付 2.微信支付;type:tinyint"`
	PayTime     utils.JSONTime `json:"payTime" form:"payTime" gorm:"column:pay_time;comment:支付時間;type:datetime"`
	OrderStatus int            `json:"orderStatus" form:"orderStatus" gorm:"column:order_status;comment:訂單狀態:0.待支付 1.已支付 2.配貨完成 3:出庫成功 4.交易成功 -1.手動關閉 -2.超時關閉 -3.商家關閉;type:tinyint"`
	ExtraInfo   string         `json:"extraInfo" form:"extraInfo" gorm:"column:extra_info;comment:訂單body;type:varchar(100);"`
	IsDeleted   int            `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:刪除標識字段(0-未刪除 1-已刪除);type:tinyint"`
	CreateTime  utils.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:創建時間;type:datetime"`
	UpdateTime  utils.JSONTime `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:最新修改時間;type:datetime"`
}

func (Order) TableName() string {
	return "tb_newbee_mall_order"
}
