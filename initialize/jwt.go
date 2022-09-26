package initialize

import (
	"bullet-screen/common"
	"bullet-screen/model/public"
)

func JwtInit() *public.Jwt {

	jwtKey := common.Global_Viper.GetString("jwt.jwt_key")
	bufferTime := common.Global_Viper.GetInt64("jwt.buffer_time")
	expiresTime := common.Global_Viper.GetInt64("jwt.expires_time")
	issuer := common.Global_Viper.GetString("jwt.issuer")

	j := &public.Jwt{
		JwtKey:      jwtKey,
		ExpiresTime: bufferTime,
		BufferTime:  expiresTime,
		Issuer:      issuer,
	}

	return j
}