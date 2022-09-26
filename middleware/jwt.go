package middleware

import (
	"bullet-screen/common"
	"bullet-screen/util"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)


func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 x-token 头部信息，登录时返回token信息
		token := c.Request.Header.Get("x-token")
		if token == "" {
			util.ResponseMap(util.TOKEN_IS_NULL,nil, c)
			c.Abort()
			return
		}

		// if token在黑名单...
		// .......

		j := util.NewJWT()

		// 解析token信息
		claims, code := j.ParseToken(token)

		if code != util.OK {
			util.ResponseMap(code, nil, c)
			c.Abort()
			return
		}

		// 校验时间
		if claims.ExpiresAt - time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + common.Global_Jwt.ExpiresTime
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		}

		c.Set("claims", claims)
		c.Next()
	}
}