package utils

import (
	"os"

	"github.com/baaj2109/newbee_mall/global"
	"go.uber.org/zap"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {
	for _, dir := range dirs {
		exist, err := PathExists(dir)
		if err != nil {
			return err
		}
		if !exist {
			global.GVA_LOG.Debug("create directory: " + dir)
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				global.GVA_LOG.Error("create directory"+dir, zap.Any(" error:", err))
				return err
			}
		}
	}
	return err
}
