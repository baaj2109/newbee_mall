package newbeemall

import (
	"github.com/baaj2109/newbee_mall/core"
	"github.com/baaj2109/newbee_mall/global"
	"github.com/baaj2109/newbee_mall/initialize"
)

func main() {

	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()
	global.GVA_DB = initialize.Gorm()
	core.RunWindowsServer()

}
