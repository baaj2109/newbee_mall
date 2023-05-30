package model

import "time"

type UserToken struct {
	UserId     int       `json:"userId" form:"userId" gorm:"primarykey;AUTO_INCREMENT"`
	Token      string    `json:"token" form:"token" gorm:"column:token;comment:token值(32位字符串);type:varchar(32);"`
	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改時間;type:datetime"`
	ExpireTime time.Time `json:"expireTime" form:"expireTime" gorm:"column:expire_time;comment:token過期時間;type:datetime"`
}

// TableName MallUserToken 表名
func (UserToken) TableName() string {
	return "tb_newbee_mall_user_token"
}
