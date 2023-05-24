package mall

import (
	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model"
)

type MallUserTokenService struct {
}

func (m *MallUserTokenService) ExistUserToken(token string) (mallUserToken model.UserToken, err error) {
	err = global.GVA_DB.Where("token =?", token).First(&mallUserToken).Error
	return
}

func (m *MallUserTokenService) DeleteMallUserToken(token string) (err error) {
	err = global.GVA_DB.Delete(&[]model.UserToken{}, "token =?", token).Error
	return err
}
