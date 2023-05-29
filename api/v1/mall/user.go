package mall

import (
	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model/request/update_request"
	"github.com/baaj2109/newbee_mall/model/request/user_request"
	"github.com/baaj2109/newbee_mall/model/response"
	"github.com/baaj2109/newbee_mall/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MallUserApi struct {
}

// / register user using gin context
func (m *MallUserApi) UserRegister(c *gin.Context) {
	var request user_request.RegisterUserParam
	_ = c.ShouldBindJSON(&request)
	if err := utils.Verify(request, utils.MallUserRegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := mallUserService.RegisterUser(request); err != nil {
		global.GVA_LOG.Error("register user failed", zap.Error(err))
		response.FailWithMessage("register user failed", c)
	}
	response.OkWithMessage("register user success", c)
}

// / update user info using gin context
func (m *MallUserApi) UserInfoUpdate(c *gin.Context) {
	var request update_request.UpdateUserInfoParam
	token := c.GetHeader("token")
	if err := mallUserService.UpdateUserInfo(token, request); err != nil {
		global.GVA_LOG.Error("update user info failed", zap.Error(err))
		response.FailWithMessage("update user info failed", c)
	}
	response.OkWithMessage("update user info success", c)
}

// / get user info using gin context
func (m *MallUserApi) GetUserInfo(c *gin.Context) {
	token := c.GetHeader("token")
	res, err := mallUserService.GetUserInfo(token)
	if err != nil {
		global.GVA_LOG.Error("get user info failed", zap.Error(err))
		response.FailWithMessage("get user info failed", c)
	}
	response.OkWithData(res, c)
}

// / login using gin context
func (m *MallUserApi) UserLogin(c *gin.Context) {
	var request user_request.UserLoginParam
	_ = c.ShouldBindJSON(&request)
	_, token, err := mallUserService.UserLogin(request)
	if err != nil {
		global.GVA_LOG.Error("logout failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	}
	response.OkWithData(token.Token, c)
}

// / logout using gin context
func (m *MallUserApi) UserLogout(c *gin.Context) {
	token := c.GetHeader("token")
	if err := mallUserTokenService.DeleteMallUserToken(token); err != nil {
		response.FailWithMessage("logout failed", c)
	}
	response.OkWithMessage("logout success", c)
}
