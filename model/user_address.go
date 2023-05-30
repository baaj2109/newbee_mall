package model

import "github.com/baaj2109/newbee_mall/utils"

type UserAddress struct {
	AddressId     int            `json:"addressId" form:"addressId" gorm:"primarykey;AUTO_INCREMENT"`
	UserId        int            `json:"userId" form:"userId" gorm:"column:user_id;comment:用戶主键id;type:bigint"`
	UserName      string         `json:"userName" form:"userName" gorm:"column:user_name;comment:收貨人姓名;type:varchar(30);"`
	UserPhone     string         `json:"userPhone" form:"userPhone" gorm:"column:user_phone;comment:收貨人手机號;type:varchar(11);"`
	DefaultFlag   int            `json:"defaultFlag" form:"defaultFlag" gorm:"column:default_flag;comment:是否為默認 0-非默認 1-是默認;type:tinyint"`
	ProvinceName  string         `json:"provinceName" form:"provinceName" gorm:"column:province_name;comment:省;type:varchar(32);"`
	CityName      string         `json:"cityName" form:"cityName" gorm:"column:city_name;comment:城;type:varchar(32);"`
	RegionName    string         `json:"regionName" form:"regionName" gorm:"column:region_name;comment:区;type:varchar(32);"`
	DetailAddress string         `json:"detailAddress" form:"detailAddress" gorm:"column:detail_address;comment:收件詳细地址(街道/樓宇/單元);type:varchar(64);"`
	IsDeleted     int            `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:刪除標識字段(0-未刪除 1-已刪除);type:tinyint"`
	CreateTime    utils.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:添加時間;type:datetime"`
	UpdateTime    utils.JSONTime `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改時間;type:datetime"`
}

// TableName MallUserAddress 表名
func (UserAddress) TableName() string {
	return "tb_newbee_mall_user_address"
}
