package controller

import (
	"bullet-screen/service"
	"bullet-screen/util"
	"github.com/gin-gonic/gin"
)

// 获取验证码
func GetCaptcha(c *gin.Context) {
	uuid, b64s, code := service.GetCaptcha()

	data := map[string]interface{} {
		"uuid": uuid,
		"b64s": b64s,
	}

	util.ResponseMap(code, data, c)
}

// 校验验证码
func VerifyCaptcha(c *gin.Context) {
	uuid := c.PostForm("uuid")
	value := c.PostForm("value")
	ans, code := service.Verify(uuid, value)

	util.ResponseData(code, ans, c)
}