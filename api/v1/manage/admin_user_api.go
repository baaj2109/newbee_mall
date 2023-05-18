package manage

import (
	"strconv"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model"
	"github.com/baaj2109/newbee_mall/model/common/request"
	"github.com/baaj2109/newbee_mall/model/common/response"
	"github.com/baaj2109/newbee_mall/model/request/update_request"
	"github.com/baaj2109/newbee_mall/model/request/user_request"
	"github.com/baaj2109/newbee_mall/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AdminUserApi struct {
}

// create admin user api by gin context
func (a *AdminUserApi) CreateAdminUser(c *gin.Context) {
	var params user_request.AdminParam
	_ = c.ShouldBindJSON(&params)
	if err := utils.Verify(params, utils.AdminUserRegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mallAdminUser := model.AdminUser{
		LoginUserName: params.LoginUserName,
		NickName:      params.NickName,
		LoginPassword: utils.MD5([]byte(params.LoginPassword)),
	}
	if err := mallAdminUserService.CreateAdminUser(mallAdminUser); err != nil {
		global.GVA_LOG.Error("create admin user failed:", zap.Error(err))
		response.FailWithMessage("create admin user failed:"+err.Error(), c)
	} else {
		response.OkWithMessage("create admin user success", c)
	}
}

// update admin user password by gin context
func (a *AdminUserApi) UpdateAdminUserPassword(c *gin.Context) {
	var params update_request.UpdatePasswordParam
	_ = c.ShouldBindJSON(&params)
	userToken := c.GetHeader("token")
	if err := mallAdminUserService.UpdateAdminUserPassword(userToken, params); err != nil {
		global.GVA_LOG.Error("update admin user password failed:", zap.Error(err))
		response.FailWithMessage("update admin user password failed:"+err.Error(), c)
	} else {
		response.OkWithMessage("update admin user password success", c)
	}
}

// update admin user name by gin context
func (a *AdminUserApi) UpdateAdminUserName(c *gin.Context) {
	var params update_request.UpdateNameParam
	_ = c.ShouldBindJSON(&params)
	userToken := c.GetHeader("token")
	if err := mallAdminUserService.UpdateAdminUserName(userToken, params); err != nil {
		global.GVA_LOG.Error("update admin user name failed:", zap.Error(err))
		response.FailWithMessage("update admin user name failed:", c)
	} else {
		response.OkWithMessage("update admin user name success", c)
	}
}

// get admin user profile by gin context
func (a *AdminUserApi) GetAdminUser(c *gin.Context) {
	adminToken := c.GetHeader("token")
	if mallAdminUser, err := mallAdminUserService.GetAdminUser(adminToken); err != nil {
		global.GVA_LOG.Error("get admin user failed:", zap.Error(err))
		response.FailWithMessage("get admin user failed:", c)
	} else {
		mallAdminUser.LoginPassword = "******"
		response.OkWithData(mallAdminUser, c)
	}
}

// admin user login by gin context
func (a *AdminUserApi) AdminUserLogin(c *gin.Context) {
	var params user_request.AdminLoginParam
	_ = c.ShouldBindJSON(&params)
	if _, adminToken, err := mallAdminUserService.AdminLogin(params); err != nil {
		response.FailWithMessage("admin user login failed", c)
	} else {
		response.OkWithData(adminToken.Token, c)
	}
}

// admin user logout by gin context
func (m *AdminUserApi) AdminUserLogout(c *gin.Context) {
	token := c.GetHeader("token")
	if err := mallAdminUserTokenService.DeleteMallAdminUserToken(token); err != nil {
		response.FailWithMessage("admin user logout failed", c)
	} else {
		response.OkWithMessage("admin user logout success", c)
	}

}

// / get user list by gin context
func (m *AdminUserApi) GetAdminUserList(c *gin.Context) {

}

// / lock user by gin context
func (m *AdminUserApi) LockUser(c *gin.Context) {
	lockStatus, _ := strconv.Atoi(c.Param("lockStatus"))
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := mallUserService.LockUser(IDS, lockStatus); err != nil {
		global.GVA_LOG.Error("update user failed!", zap.Error(err))
		response.FailWithMessage("update user failed!", c)
	} else {
		response.OkWithMessage("update user success", c)
	}
}
