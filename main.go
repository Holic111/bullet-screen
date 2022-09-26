package main

import (
	"bullet-screen/common"
	"bullet-screen/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	common.Global_Viper = initialize.ViperInit()
	common.Global_AppMode = initialize.AppModeInit()
	common.Global_Mysql = initialize.MySQLInit()
	common.Global_Redis = initialize.RedisInit()
	common.Global_Jwt = initialize.JwtInit()
	common.Global_Zap = initialize.ZapInit()


	initialize.RouterInit()
}