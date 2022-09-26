package initialize

import (
	"bullet-screen/common"
	"bullet-screen/model/public"
)

func AppModeInit() *public.AppMode {

	mode := common.Global_Viper.GetString("appmode.mode")
	port := common.Global_Viper.GetString("appmode.port")

	return &public.AppMode{Mode: mode, Port: port}
}