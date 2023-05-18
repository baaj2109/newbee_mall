package manage

import (
	"errors"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model"
	"github.com/baaj2109/newbee_mall/model/request/update_request"
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
