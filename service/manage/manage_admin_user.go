package manage

import (
	"errors"
	"time"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model"
	"github.com/baaj2109/newbee_mall/model/request/update_request"
	"github.com/baaj2109/newbee_mall/model/request/user_request"
	"github.com/baaj2109/newbee_mall/utils"
	"gorm.io/gorm"
)

type ManageAdminUserService struct {
}

func (m *ManageAdminUserService) CreateAdminUser(adminUser model.AdminUser) error {
	if !errors.Is(global.GVA_DB.Where("login_user_name = ?", adminUser.LoginUserName).First(&model.AdminUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("user already exists")
	}
	return global.GVA_DB.Create(&adminUser).Error
}

func (m *ManageAdminUserService) UpdateAdminUserName(token string, request update_request.UpdateNameParam) error {
	var adminUserToken model.AdminUserToken
	err := global.GVA_DB.Where("token = ?", token).First(&adminUserToken).Error
	if err != nil {
		return errors.New("token not found")
	}
	return global.GVA_DB.Where("admin_user_id = ?", adminUserToken.AdminUserId).Updates(
		&model.AdminUser{
			LoginUserName: request.LoginUserName,
			NickName:      request.NickName,
		}).Error
}

func (m *ManageAdminUserService) UpdateAdminUserPassword(token string, request update_request.UpdatePasswordParam) error {
	var adminUserToken model.AdminUserToken
	err := global.GVA_DB.Where("token = ?", token).First(&adminUserToken).Error
	if err != nil {
		return errors.New("token not found")
	}
	var adminUser model.AdminUser
	err = global.GVA_DB.Where("admin_user_id = ?", adminUserToken.AdminUserId).First(&adminUser).Error
	if err != nil {
		return errors.New("user not found")
	}
	if adminUser.LoginPassword != request.OriginalPassword {
		return errors.New("wrong password")
	}
	adminUser.LoginPassword = request.NewPassword

	return global.GVA_DB.Where("admin_user_id = ?", adminUserToken.AdminUserId).Updates(adminUser).Error
}

func (m *ManageAdminUserService) GetAdminUser(token string) (model.AdminUser, error) {
	var adminUserToken model.AdminUserToken
	err := global.GVA_DB.Where("token = ?", token).First(&adminUserToken).Error
	if err != nil {
		return model.AdminUser{}, errors.New("token not found")
	}
	var adminUser model.AdminUser
	err = global.GVA_DB.Where("admin_user_id = ?", adminUserToken.AdminUserId).First(&adminUser).Error
	return adminUser, err
}

func (m *ManageAdminUserService) AdminLogin(params user_request.AdminLoginParam) (mallAdminUser model.AdminUser, adminToken model.AdminUserToken, err error) {
	err = global.GVA_DB.Where("login_user_name=? AND login_password=?", params.UserName, params.PasswordMd5).First(&mallAdminUser).Error
	if mallAdminUser != (model.AdminUser{}) {
		token := utils.GetNewToken(time.Now().UnixNano()/1e6, int(mallAdminUser.AdminUserId))
		global.GVA_DB.Where("admin_user_id", mallAdminUser.AdminUserId).First(&adminToken)
		nowDate := time.Now()
		// token expired after 48h
		expireTime, _ := time.ParseDuration("48h")
		expireDate := nowDate.Add(expireTime)
		// if there is no token create, then update
		if adminToken == (model.AdminUserToken{}) {
			adminToken.AdminUserId = mallAdminUser.AdminUserId
			adminToken.Token = token
			adminToken.UpdateTime = nowDate
			adminToken.ExpireTime = expireDate
			if err = global.GVA_DB.Create(&adminToken).Error; err != nil {
				return
			}
		} else {
			adminToken.Token = token
			adminToken.UpdateTime = nowDate
			adminToken.ExpireTime = expireDate
			if err = global.GVA_DB.Save(&adminToken).Error; err != nil {
				return
			}
		}
	}
	return mallAdminUser, adminToken, err

}
