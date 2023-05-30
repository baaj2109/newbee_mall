package mall

import (
	v1 "github.com/baaj2109/newbee_mall/api/v1"
	"github.com/baaj2109/newbee_mall/middleware"
	"github.com/gin-gonic/gin"
)

type MallUserAddressRouter struct {
}

func (m *MallUserAddressRouter) InitMallUserAddressRouter(Router *gin.RouterGroup) {
	mallUserAddressRouter := Router.Group("v1").Use(middleware.UserJWTAuth())
	var mallUserAddressApi = v1.ApiGroupApp.MallApiGroup.MallUserAddressApi
	{
		mallUserAddressRouter.GET("/address", mallUserAddressApi.GetAddressList)                    //用户地址
		mallUserAddressRouter.POST("/address", mallUserAddressApi.SaveAddress)                      //添加地址
		mallUserAddressRouter.PUT("/address", mallUserAddressApi.UpdateAddress)                     //修改用户地址
		mallUserAddressRouter.GET("/address/:addressId", mallUserAddressApi.GetAddress)             //获取地址详情
		mallUserAddressRouter.GET("/address/default", mallUserAddressApi.GetMallUserDefaultAddress) //获取默认地址
		mallUserAddressRouter.DELETE("/address/:addressId", mallUserAddressApi.DeleteAddress)       //删除地址
	}
}
