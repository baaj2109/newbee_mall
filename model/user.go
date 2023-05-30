package model

import "github.com/baaj2109/newbee_mall/utils"

type User struct {
	UserId        int            `json:"userId" form:"userId" gorm:"primarykey;AUTO_INCREMENT"`
	NickName      string         `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:用戶昵稱;type:varchar(50);"`
	LoginName     string         `json:"loginName" form:"loginName" gorm:"column:login_name;comment:登入名稱(默認為手机號);type:varchar(11);"`
	PasswordMd5   string         `json:"passwordMd5" form:"passwordMd5" gorm:"column:password_md5;comment:MD5加密後的密碼;type:varchar(32);"`
	IntroduceSign string         `json:"introduceSign" form:"introduceSign" gorm:"column:introduce_sign;comment:個性簽名;type:varchar(100);"`
	IsDeleted     int            `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:注銷標識字段(0-正常 1-已注銷);type:tinyint"`
	LockedFlag    int            `json:"lockedFlag" form:"lockedFlag" gorm:"column:locked_flag;comment:鎖定標識字段(0-未鎖定 1-已鎖定);type:tinyint"`
	CreateTime    utils.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:註冊時間;type:datetime"`
}

// TableName MallUser 表名
func (User) TableName() string {
	return "tb_newbee_mall_user"
}
