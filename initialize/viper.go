package initialize

import (
	"fmt"
	"github.com/spf13/viper"
)

func ViperInit() *viper.Viper{
	v := viper.New()
	//建立默认值
	//v.SetDefault("ContentDir", "content")

	v.SetConfigName("conf")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config/")

	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置文件错误", err)
	}
	return v
}