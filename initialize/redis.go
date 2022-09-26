package initialize

import (
	"bullet-screen/common"
	"github.com/go-redis/redis"
)

func RedisInit() *redis.Client {
	address := common.Global_Viper.GetString("redis.address")
	password := common.Global_Viper.GetString("redis.password")
	db := common.Global_Viper.GetInt("redis.db")
	poolsize := common.Global_Viper.GetInt("redis.poolsize")

	rdb := redis.NewClient(&redis.Options{
		Addr: address,
		Password: password,
		DB: db,
		PoolSize: poolsize,
	})

	return rdb
}