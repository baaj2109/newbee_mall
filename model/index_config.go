package model

import "github.com/baaj2109/newbee_mall/utils"

type IndexConfig struct {
	ConfigId    int            `json:"configId" form:"configId" gorm:"primarykey;AUTO_INCREMENT"`
	ConfigName  string         `json:"configName" form:"configName" gorm:"column:config_name;comment:顯示字符(配置搜索時不可為空，其他可為空);type:varchar(50);"`
	ConfigType  int            `json:"configType" form:"configType" gorm:"column:config_type;comment:1-搜索框熱搜 2-搜索下拉框熱搜 3-(首頁)熱銷商品 4-(首頁)新品上線 5-(首頁)為你推薦;type:tinyint"`
	GoodsId     int            `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:商品id 默認為0;type:bigint"`
	RedirectUrl string         `json:"redirectUrl" form:"redirectUrl" gorm:"column:redirect_url;comment:點擊後的跳轉地址(默認不跳轉);type:varchar(100);"`
	ConfigRank  int            `json:"configRank" form:"configRank" gorm:"column:config_rank;comment:排序值(字段越大越靠前);type:int"`
	IsDeleted   int            `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:刪除標識字段(0-未刪除 1-已刪除);type:tinyint"`
	CreateTime  utils.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:創建時間;type:datetime"`
	CreateUser  int            `json:"createUser" form:"createUser" gorm:"column:create_user;comment:創建者id;type:int"`
	UpdateTime  utils.JSONTime `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:最新修改時間;type:datetime"`
	UpdateUser  int            `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改者id;type:int"`
}

// TableName MallIndexConfig 表名
func (IndexConfig) TableName() string {
	return "tb_newbee_mall_index_config"
}
