package routers

import (
	"bullet-screen/controller"
	"github.com/gin-gonic/gin"
)

type CaptchaRouter struct {}

func (capt *CaptchaRouter) CaptchaRouters(r *gin.RouterGroup) {
	c := r.Group("/captcha")
	{
		// 获取验证码
		c.GET("/get", controller.GetCaptcha)

		// 校验验证码
		c.POST("/verify", controller.VerifyCaptcha)
	}
}