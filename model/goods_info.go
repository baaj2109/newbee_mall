package model

import "github.com/baaj2109/newbee_mall/utils"

type GoodsInfo struct {
	GoodsId            int            `json:"goodsId" form:"goodsId" gorm:"primarykey;AUTO_INCREMENT"`
	GoodsName          string         `json:"goodsName" form:"goodsName" gorm:"column:goods_name;comment:商品名;type:varchar(200);"`
	GoodsIntro         string         `json:"goodsIntro" form:"goodsIntro" gorm:"column:goods_intro;comment:商品簡介;type:varchar(200);"`
	GoodsCategoryId    int            `json:"goodsCategoryId" form:"goodsCategoryId" gorm:"column:goods_category_id;comment:關聯分類id;type:bigint"`
	GoodsCoverImg      string         `json:"goodsCoverImg" form:"goodsCoverImg" gorm:"column:goods_cover_img;comment:商品主圖;type:varchar(200);"`
	GoodsCarousel      string         `json:"goodsCarousel" form:"goodsCarousel" gorm:"column:goods_carousel;comment:商品輪播圖;type:varchar(500);"`
	GoodsDetailContent string         `json:"goodsDetailContent" form:"goodsDetailContent" gorm:"column:goods_detail_content;comment:商品詳情;type:text;"`
	OriginalPrice      int            `json:"originalPrice" form:"originalPrice" gorm:"column:original_price;comment:商品價格;type:int"`
	SellingPrice       int            `json:"sellingPrice" form:"sellingPrice" gorm:"column:selling_price;comment:商品實際售價;type:int"`
	StockNum           int            `json:"stockNum" form:"stockNum" gorm:"column:stock_num;comment:商品庫存數量;type:int"`
	Tag                string         `json:"tag" form:"tag" gorm:"column:tag;comment:商品標籤;type:varchar(20);"`
	GoodsSellStatus    int            `json:"goodsSellStatus" form:"goodsSellStatus" gorm:"column:goods_sell_status;comment:商品上架狀態 1-下架 0-上架;type:tinyint"`
	CreateUser         int            `json:"createUser" form:"createUser" gorm:"column:create_user;comment:添加者主键id;type:int"`
	CreateTime         utils.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:商品添加時間;type:datetime"`
	UpdateUser         int            `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改者主键id;type:int"`
	UpdateTime         utils.JSONTime `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:商品修改時間;type:datetime"`
}

// TableName MallGoodsInfo 表名
func (GoodsInfo) TableName() string {
	return "tb_newbee_mall_goods_info"
}
