package model

import "github.com/baaj2109/newbee_mall/utils"

type Carousel struct {
	CarouselId   int            `json:"carouselId" form:"carouselId" gorm:"primarykey;AUTO_INCREMENT"`
	CarouselUrl  string         `json:"carouselUrl" form:"carouselUrl" gorm:"column:carousel_url;comment:輪播圖;type:varchar(100);"`
	RedirectUrl  string         `json:"redirectUrl" form:"redirectUrl" gorm:"column:redirect_url;comment:點擊後的跳轉地址(默認不跳轉);type:varchar(100);"`
	CarouselRank int            `json:"carouselRank" form:"carouselRank" gorm:"column:carousel_rank;comment:排序值(字段越大越靠前);type:int"`
	IsDeleted    int            `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:刪除標識字段(0-未刪除 1-已刪除);type:tinyint"`
	CreateTime   utils.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:創建時間;type:datetime"`
	CreateUser   int            `json:"createUser" form:"createUser" gorm:"column:create_user;comment:創建者id;type:int"`
	UpdateTime   utils.JSONTime `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改時間;type:datetime"`
	UpdateUser   int            `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改者id;type:int"`
}

func (Carousel) TableName() string {
	return "tb_newbee_mall_carousel"
}
