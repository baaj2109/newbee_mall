package manage

import (
	v1 "github.com/baaj2109/newbee_mall/api/v1"
	"github.com/baaj2109/newbee_mall/middleware"
	"github.com/gin-gonic/gin"
)

type ManageAdminUserRouter struct {
}

func (r *ManageAdminUserRouter) InitManageAdminUserRouter(Router *gin.RouterGroup) {
	mallAdminUserRouter := Router.Group("v1").Use(middleware.AdminJWTAuth())
	mallAdminUserWithoutRouter := Router.Group("v1")
	var mallAdminUserApi = v1.ApiGroupApp.ManageApiGroup.AdminUserApi
	{
		mallAdminUserRouter.POST("createMallAdminUser", mallAdminUserApi.CreateAdminUser) // 新建MallAdminUser
		mallAdminUserRouter.PUT("adminUser/name", mallAdminUserApi.UpdateAdminUserName)   // 更新MallAdminUser
		mallAdminUserRouter.PUT("adminUser/password", mallAdminUserApi.UpdateAdminUserPassword)
		// mallAdminUserRouter.GET("users", mallAdminUserApi.UserList)
		mallAdminUserRouter.PUT("users/:lockStatus", mallAdminUserApi.LockUser)
		mallAdminUserRouter.GET("adminUser/profile", mallAdminUserApi.GetAdminUser) // 根據ID獲取 admin詳情
		mallAdminUserRouter.DELETE("logout", mallAdminUserApi.AdminUserLogout)
		// 	mallAdminUserRouter.POST("upload/file", mallAdminUserApi.UploadFile) //上傳圖片

	}
	{
		mallAdminUserWithoutRouter.POST("adminUser/login", mallAdminUserApi.AdminUserLogin) //管理員登入
	}
}
