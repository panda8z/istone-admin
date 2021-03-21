package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/panda8z/istone/common"
	"github.com/panda8z/istone/global"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	fmt.Println(path)
	// TODO: 0001: Viper 方法可以从命令行，环境变量，默认配置文件 获取参数
	// 目前仅实现默认配置文件
	v := viper.New()
	v.SetConfigFile(common.ConfigDefaultPath)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("config file 读取出错: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file 已经更新:", e.Name)
		unmarshal(v)
	})
	unmarshal(v)
	return v
}

func unmarshal(v *viper.Viper) {
	if err := v.Unmarshal(&global.ISA_CONFIG); err != nil {
		fmt.Println(err)
	}
}
