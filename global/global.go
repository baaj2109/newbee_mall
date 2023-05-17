package global

import (
	"github.com/baaj2109/newbee_mall/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	GVA_LOG    *zap.Logger
)
