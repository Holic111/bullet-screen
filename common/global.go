package common

import (
	"bullet-screen/model/public"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Global_Viper *viper.Viper
	Global_Mysql *gorm.DB
	Global_Redis *redis.Client
	Global_AppMode *public.AppMode
	Global_Jwt     *public.Jwt

	Global_Zap *zap.Logger
)