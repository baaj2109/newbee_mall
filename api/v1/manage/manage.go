package manage

import "github.com/baaj2109/newbee_mall/service"

type ManageGroup struct {
	AdminUserApi
	ManageGoodsCategoryApi
	// ManageGoodsInfoApi
	ManageCarouselApi
	// ManageIndexConfigApi
	// ManageOrderApi
}

var mallAdminUserService = service.ServiceGroupApp.ManageServiceGroup.ManageAdminUserService
var mallAdminUserTokenService = service.ServiceGroupApp.ManageServiceGroup.ManageAdminUserTokenService
var mallUserService = service.ServiceGroupApp.ManageServiceGroup.UserStatusService
var mallGoodsCategoryService = service.ServiceGroupApp.ManageServiceGroup.ManageGoodsCategoryService

// var fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
var mallGoodsInfoService = service.ServiceGroupApp.ManageServiceGroup.ManageGoodsInfoService
var mallCarouselService = service.ServiceGroupApp.ManageServiceGroup.ManageCarouselService
var mallIndexConfigService = service.ServiceGroupApp.ManageServiceGroup.ManageIndexConfigService
var mallOrderService = service.ServiceGroupApp.ManageServiceGroup.ManageOrderService
