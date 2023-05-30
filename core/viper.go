package core

import (
	"flag"
	"fmt"
	"os"

	"github.com/baaj2109/newbee_mall/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {

	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 優先級: 命令行 > 環境變量 > 默認值
			if configEnv := os.Getenv("BEEMALL_CONFIG"); configEnv == "" {
				config = "config.yaml"
				fmt.Printf("[viper] use config.yaml")
			} else {
				config = configEnv
				fmt.Printf("[viper] use config in %s\n", config)
			}
		} else {

			fmt.Print("[viper] use config with -c command in " + config + "\n")
		}

	} else {
		config = path[0]
		fmt.Printf("[viper] use config in %s\n", config)
	}
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
