package newbeemall

import (
	"github.com/baaj2109/newbee_mall/core"
	"github.com/baaj2109/newbee_mall/global"
)

func main() {

	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()

}
