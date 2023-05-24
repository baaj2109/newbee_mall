package mall

import (
	"errors"
	"time"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/model"
	"github.com/baaj2109/newbee_mall/model/request/update_request"
	"github.com/baaj2109/newbee_mall/model/request/user_request"
	"github.com/baaj2109/newbee_mall/model/response"
	"github.com/baaj2109/newbee_mall/utils"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type MallUserService struct {
}

// / register user
func (m *MallUserService) RegisterUser(req user_request.RegisterUserParam) (err error) {
	if !errors.Is(global.GVA_DB.Where("login_name = ?", req.LoginName).First(&model.User{}).Error,
		gorm.ErrRecordNotFound) {
		return errors.New("user already exists")
	}
	return global.GVA_DB.Create(&model.User{
		LoginName:     req.LoginName,
		PasswordMd5:   utils.MD5([]byte(req.Password)),
		IntroduceSign: "no sign",
		CreateTime:    utils.JSONTime{Time: time.Now()},
	}).Error
}

// / update user info
func (m *MallUserService) UpdateUserInfo(token string, req update_request.UpdateUserInfoParam) (err error) {
	var userToken model.UserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("user can't find")
	}
	var user model.User

	err = global.GVA_DB.Where("user_id =?", userToken.UserId).First(&user).Error
	if err != nil {
		return errors.New("user id can't find")
	}
	/// if password is empty, means user doesn't want to modify password, use origin password
	if !(req.PasswordMd5 == "") {
		user.PasswordMd5 = utils.MD5([]byte(req.PasswordMd5))
	}
	user.NickName = req.NickName
	user.IntroduceSign = req.IntroduceSign
	err = global.GVA_DB.Where("user_id =?", userToken.UserId).UpdateColumns(&user).Error
	return
}

func (m *MallUserService) GetUserInfo(token string) (res response.MallUserDetailResponse, err error) {
	var userToken model.UserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return res, errors.New("user can't find")
	}
	var userInfo model.User
	err = global.GVA_DB.Where("user_id =?", userToken.UserId).First(&userInfo).Error
	if err != nil {
		return res, errors.New("failed to get user info")
	}
	err = copier.Copy(&res, &userInfo)
	return
}

func (m *MallUserService) UserLogin(params user_request.UserLoginParam) (user model.User, userToken model.UserToken, err error) {
	err = global.GVA_DB.Where("login_name=? AND password_md5=?", params.LoginName, params.PasswordMd5).First(&user).Error
	if user != (model.User{}) {
		token := utils.GetNewToken(time.Now().UnixNano()/1e6, int(user.UserId))
		global.GVA_DB.Where("user_id", user.UserId).First(&token)
		nowDate := time.Now()
		/// expired after 48 hours
		expireTime, _ := time.ParseDuration("48h")
		expireDate := nowDate.Add(expireTime)
		// update token, if token is empty
		if userToken == (model.UserToken{}) {
			userToken.UserId = user.UserId
			userToken.Token = token
			userToken.UpdateTime = nowDate
			userToken.ExpireTime = expireDate
			if err = global.GVA_DB.Save(&userToken).Error; err != nil {
				return
			}
		} else {
			userToken.Token = token
			userToken.UpdateTime = nowDate
			userToken.ExpireTime = expireDate
			if err = global.GVA_DB.Save(&userToken).Error; err != nil {
				return
			}
		}
	}
	return user, userToken, err
}
