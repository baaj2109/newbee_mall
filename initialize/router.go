package initialize

import (
	"github.com/baaj2109/newbee_mall/router"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	return router.GetRouter()
}
