package middleware

import (
	"time"

	"github.com/baaj2109/newbee_mall/model/response"
	"github.com/baaj2109/newbee_mall/service"
	"github.com/gin-gonic/gin"
)

var manageAdminUserTokenService = service.ServiceGroupApp.ManageServiceGroup.ManageAdminUserTokenService

// var mallUserTokenService = service.ServiceGroupApp.MallServiceGroup.MallUserTokenService

func AdminJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.FailWithDetailed(nil, "not login or invalid token", c)
			c.Abort()
			return
		}
		mallAdminUserToken, err := manageAdminUserTokenService.IsAdminTokenExist(token)
		if err != nil {
			response.FailWithDetailed(nil, "not login or invalid token", c)
			c.Abort()
			return
		}
		if time.Now().After(mallAdminUserToken.ExpireTime) {
			response.FailWithDetailed(nil, "token expired", c)
			err = manageAdminUserTokenService.DeleteMallAdminUserToken(token)
			if err != nil {
				return
			}
			c.Abort()
			return
		}
		c.Next()
	}
}

// func UserJWTAuth() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		token := c.Request.Header.Get("token")
// 		if token == "" {
// 			response.UnLogin(nil, c)
// 			c.Abort()
// 			return
// 		}
// 		err, mallUserToken := mallUserTokenService.IsUserTokenExist(token)
// 		if err != nil {
// 			response.UnLogin(nil, c)
// 			c.Abort()
// 			return
// 		}
// 		if time.Now().After(mallUserToken.ExpireTime) {
// 			response.FailWithDetailed(nil, "token expired", c)
// 			err = mallUserTokenService.DeleteMallUserToken(token)
// 			if err != nil {
// 				return
// 			}
// 			c.Abort()
// 			return
// 		}
// 		c.Next()
// 	}

// }
