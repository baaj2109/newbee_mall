package utils

import (
	"os"

	"github.com/natefinch/lumberjack"

	"github.com/baaj2109/newbee_mall/global"
	"go.uber.org/zap/zapcore"
)

func GetWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, //日誌位置
		MaxSize:    10,   //在切割之前，日誌的最大大小（MB）
		MaxBackups: 200,  //保留文件的最大个数
		MaxAge:     30,   //保留文件的最大天数
		Compress:   true, //是否壓縮文件
	}

	if global.GVA_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}
