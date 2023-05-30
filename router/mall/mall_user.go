package mall

import (
	v1 "github.com/baaj2109/newbee_mall/api/v1"
	"github.com/baaj2109/newbee_mall/middleware"
	"github.com/gin-gonic/gin"
)

type MallUserRouter struct {
}

func (m *MallUserRouter) InitMallUserRouter(Router *gin.RouterGroup) {
	mallUserRouter := Router.Group("v1").Use(middleware.UserJWTAuth())
	userRouter := Router.Group("v1")
	var mallUserApi = v1.ApiGroupApp.MallApiGroup.MallUserApi
	{
		mallUserRouter.PUT("/user/info", mallUserApi.UserInfoUpdate) //修改用戶訊息
		mallUserRouter.GET("/user/info", mallUserApi.GetUserInfo)    //獲取用戶訊息
		mallUserRouter.POST("/user/logout", mallUserApi.UserLogout)  //登出
	}
	{
		userRouter.POST("/user/register", mallUserApi.UserRegister) //用戶註冊
		userRouter.POST("/user/login", mallUserApi.UserLogin)       //登入

	}

}
