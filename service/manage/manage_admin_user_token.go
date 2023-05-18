package manage

import (
	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model"
)

type ManageAdminUserTokenService struct {
}

func (m *ManageAdminUserTokenService) IsAdminTokenExist(token string) (mallAdminUserToken model.AdminUserToken, err error) {
	err = global.GVA_DB.Where("token =?", token).First(&mallAdminUserToken).Error
	return mallAdminUserToken, err
}

func (m *ManageAdminUserTokenService) DeleteMallAdminUserToken(token string) (err error) {
	err = global.GVA_DB.Delete(&[]model.AdminUserToken{}, "token =?", token).Error
	return err
}
