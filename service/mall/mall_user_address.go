package mall

import (
	"errors"
	"time"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model"
	"github.com/baaj2109/newbee_mall/model/request/address_request"
	"github.com/baaj2109/newbee_mall/model/request/update_request"
	"github.com/baaj2109/newbee_mall/utils"
	"github.com/jinzhu/copier"
)

type MallUserAddressService struct {
}

func (m *MallUserAddressService) GetUserAddress(token string) (userAddress []model.UserAddress, err error) {
	var userToken model.UserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return userAddress, errors.New("user cant find")
	}
	global.GVA_DB.Where("user_id=? and is_deleted=0", userToken.UserId).Find(&userAddress)
	return
}

// / save user address
func (m *MallUserAddressService) SaveUserAddress(token string, req address_request.AddAddressParam) (err error) {
	var userToken model.UserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("user can't find")
	}

	var defaultAddress model.UserAddress
	copier.Copy(&defaultAddress, &req)
	defaultAddress.CreateTime = utils.JSONTime{Time: time.Now()}
	defaultAddress.UpdateTime = utils.JSONTime{Time: time.Now()}
	defaultAddress.UserId = userToken.UserId
	if req.DefaultFlag == 1 {
		global.GVA_DB.Where("user_id=? and default_flag =1 and is_deleted = 0", userToken.UserId).First(&defaultAddress)
		if defaultAddress != (model.UserAddress{}) {
			defaultAddress.UpdateTime = utils.JSONTime{Time: time.Now()}
			err = global.GVA_DB.Save(&defaultAddress).Error
			if err != nil {
				return
			}
		}
	} else {
		err = global.GVA_DB.Create(&defaultAddress).Error
		if err != nil {
			return
		}
	}
	return
}

// / update user address
func (m *MallUserAddressService) UpdateUserAddress(token string, req update_request.UpdateAddressParam) (err error) {
	var userToken model.UserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("user can't find")
	}
	var userAddress model.UserAddress
	if err = global.GVA_DB.Where("address_id =? and user_id =?", req.AddressId, userToken.UserId).First(&userAddress).Error; err != nil {
		return errors.New("address can't find")
	}
	if userToken.UserId != userAddress.UserId {
		return errors.New("token id and address id must be same")
	}
	if req.DefaultFlag == 1 {
		var defaultUserAddress model.UserAddress
		global.GVA_DB.Where("user_id=? and default_flag =1 and is_deleted = 0", userToken.UserId).First(&defaultUserAddress)
		if defaultUserAddress != (model.UserAddress{}) {
			defaultUserAddress.DefaultFlag = 0
			defaultUserAddress.UpdateTime = utils.JSONTime{Time: time.Now()}
			err = global.GVA_DB.Save(&defaultUserAddress).Error
			if err != nil {
				return
			}
		}
	}
	err = copier.Copy(&userAddress, &req)
	if err != nil {
		return
	}
	userAddress.UpdateTime = utils.JSONTime{Time: time.Now()}
	userAddress.UserId = userToken.UserId
	err = global.GVA_DB.Save(&userAddress).Error
	return

}

func (m *MallUserAddressService) GetUserAddressById(token string, id int) (userAddress model.UserAddress, err error) {
	var userToken model.UserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return userAddress, errors.New("user can't find")
	}
	if err = global.GVA_DB.Where("address_id =?", id).First(&userAddress).Error; err != nil {
		return userAddress, errors.New("address can't find")
	}
	if userToken.UserId != userAddress.UserId {
		return userAddress, errors.New("token id and address id must be same")
	}
	return
}

func (m *MallUserAddressService) GetUserDefaultAddress(token string) (userAddress model.UserAddress, err error) {
	var userToken model.UserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return userAddress, errors.New("user can't find")
	}
	if err = global.GVA_DB.Where("user_id =? and default_flag =1 and is_deleted = 0 ", userToken.UserId).First(&userAddress).Error; err != nil {
		return userAddress, errors.New("default address can't find")
	}
	return
}

func (m *MallUserAddressService) DeleteUserAddress(token string, id int) (err error) {
	var userToken model.UserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("user can't find")
	}
	var userAddress model.UserAddress
	if err = global.GVA_DB.Where("address_id =?", id).First(&userAddress).Error; err != nil {
		return errors.New("address can;t find")
	}
	if userToken.UserId != userAddress.UserId {
		return errors.New("token id and address id must be same")
	}
	err = global.GVA_DB.Delete(&userAddress).Error
	return

}
