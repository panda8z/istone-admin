package main

import (
	"fmt"
	"github.com/panda8z/istone/global"
	"github.com/panda8z/istone/initialize"
)

func main() {
	// 初始化配置
	// 初始化 mysql-tables jwt zap
	global.ISA_VP = initialize.Viper()
	global.ISA_LOG = initialize.Zap()
	global.ISA_DB = initialize.Gorm()
	SayHello()
}

func SayHello() {
	fmt.Println("Hello iStone-Admin!")
}
