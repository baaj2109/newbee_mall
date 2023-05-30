package model

type AdminUser struct {
	AdminUserId   int    `json:"adminUserId" form:"adminUserId" gorm:"primarykey;AUTO_INCREMENT"`
	LoginUserName string `json:"loginUserName" form:"loginUserName" gorm:"column:login_user_name;comment:管理员登入名稱;type:varchar(50);"`
	LoginPassword string `json:"loginPassword" form:"loginPassword" gorm:"column:login_password;comment:管理员登入密碼;type:varchar(50);"`
	NickName      string `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:管理员顯示昵稱;type:varchar(50);"`
	Locked        int    `json:"locked" form:"locked" gorm:"column:locked;comment:是否鎖定 0未鎖定 1已鎖定無法登入;type:tinyint"`
}

func (AdminUser) TableName() string {
	return "tb_newbee_mall_admin_user"
}
