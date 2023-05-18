package update_request

import "github.com/baaj2109/newbee_mall/utils"

type GoodsInfoUpdateParam struct {
	GoodsId            string         `json:"goodsId"`
	GoodsName          string         `json:"goodsName"`
	GoodsIntro         string         `json:"goodsIntro"`
	GoodsCategoryId    int            `json:"goodsCategoryId"`
	GoodsCoverImg      string         `json:"goodsCoverImg"`
	GoodsCarousel      string         `json:"goodsCarousel"`
	GoodsDetailContent string         `json:"goodsDetailContent"`
	OriginalPrice      string         `json:"originalPrice"`
	SellingPrice       int            `json:"sellingPrice"`
	StockNum           string         `json:"stockNum"`
	Tag                string         `json:"tag"`
	GoodsSellStatus    int            `json:"goodsSellStatus"`
	UpdateUser         int            `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改者主键id;type:int"`
	UpdateTime         utils.JSONTime `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:商品修改时间;type:datetime"`
}
