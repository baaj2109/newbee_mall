package mall

import (
	"strconv"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model/request/address_request"
	"github.com/baaj2109/newbee_mall/model/request/update_request"
	"github.com/baaj2109/newbee_mall/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MallUserAddressApi struct{}

// / get address list using gin context
func (m *MallUserAddressApi) GetAddressList(c *gin.Context) {
	token := c.GetHeader("token")
	addressList, err := mallUserAddressService.GetUserAddress(token)
	if err != nil {
		global.GVA_LOG.Error("get address list failed", zap.Error(err))
		response.FailWithMessage("get address list failed", c)
	} else if len(addressList) < 1 {
		response.OkWithData(nil, c)
	} else {
		response.OkWithData(addressList, c)
	}
}

// / save user address using gin context
func (m *MallUserAddressApi) SaveAddress(c *gin.Context) {
	var request address_request.AddAddressParam
	_ = c.ShouldBindJSON(&request)
	token := c.GetHeader("token")
	if err := mallUserAddressService.SaveUserAddress(token, request); err != nil {
		global.GVA_LOG.Error("save user address failed", zap.Error(err))
		response.FailWithMessage("save user address failed", c)
	} else {
		response.OkWithMessage("save user address success", c)
	}
}

// / update user address using gin context
func (m *MallUserAddressApi) UpdateAddress(c *gin.Context) {
	var request update_request.UpdateAddressParam
	_ = c.ShouldBindJSON(&request)
	token := c.GetHeader("token")
	if err := mallUserAddressService.UpdateUserAddress(token, request); err != nil {
		global.GVA_LOG.Error("update user address failed", zap.Error(err))
		response.FailWithMessage("update user address failed", c)
	}
	response.OkWithMessage("update user address success", c)
}

// / get user address using gin context
func (m *MallUserAddressApi) GetAddress(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("addressId"))
	token := c.GetHeader("token")
	res, err := mallUserAddressService.GetUserAddressById(token, id)
	if err != nil {
		global.GVA_LOG.Error("get user address failed", zap.Error(err))
		response.FailWithMessage("get user address failed", c)
	} else {
		response.OkWithData(res, c)
	}
}

func (m *MallUserAddressApi) GetMallUserDefaultAddress(c *gin.Context) {
	token := c.GetHeader("token")
	res, err := mallUserAddressService.GetUserAddress(token)
	if err != nil {
		global.GVA_LOG.Error("get user address failed", zap.Error(err))
		response.FailWithMessage("get user address failed", c)
	}
	response.OkWithData(res, c)
}

func (m *MallUserAddressApi) DeleteAddress(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("addressId"))
	token := c.GetHeader("token")
	if err := mallUserAddressService.DeleteUserAddress(token, id); err != nil {
		global.GVA_LOG.Error("delete user address failed", zap.Error(err))
		response.FailWithMessage("delete user address failed", c)
	} else {
		response.OkWithMessage("delete user address success", c)
	}
}
