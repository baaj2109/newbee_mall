package model

import "github.com/baaj2109/newbee_mall/utils"

type GoodsCategory struct {
	CategoryId    int            `json:"categoryId" gorm:"primarykey;AUTO_INCREMENT"`
	CategoryLevel int            `json:"categoryLevel" gorm:"comment:分類等級"`
	ParentId      int            `json:"parentId" gorm:"comment:父類id"`
	CategoryName  string         `json:"categoryName" gorm:"comment:分類名稱"`
	CategoryRank  int            `json:"categoryRank" gorm:"comment:排序比重"`
	IsDeleted     int            `json:"isDeleted" gorm:"comment:是否刪除"`
	CreateTime    utils.JSONTime `json:"createTime" gorm:"column:create_time;comment:創建時間;type:datetime"` // 創建時間
	UpdateTime    utils.JSONTime `json:"updateTime" gorm:"column:update_time;comment:修改時間;type:datetime"` // 更新時間
}

func (GoodsCategory) TableName() string {
	return "tb_newbee_mall_goods_category"
}
