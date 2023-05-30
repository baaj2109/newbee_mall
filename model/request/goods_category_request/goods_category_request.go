package goods_category_request

import "github.com/baaj2109/newbee_mall/utils"

type GoodsCategoryReq struct {
	CategoryId    int            `json:"categoryId"`
	CategoryLevel int            `json:"categoryLevel" `
	ParentId      int            `json:"parentId"`
	CategoryName  string         `json:"categoryName" `
	CategoryRank  string         `json:"categoryRank" `
	IsDeleted     int            `json:"isDeleted" `
	CreateTime    utils.JSONTime `json:"createTime" ` // 創建時間
	UpdateTime    utils.JSONTime `json:"updateTime" ` // 更新時間
}
