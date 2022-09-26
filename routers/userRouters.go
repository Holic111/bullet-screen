package routers

import (
	"bullet-screen/controller"
	"bullet-screen/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {}

func (us *UserRouter) UserRouters(r *gin.RouterGroup) {

	u1 := r.Group("/user")
	u1.Use(middleware.JWTAuth())
	{
		// 获取所有用户信息
		u1.GET("/list", controller.GetUserList)

		// 获取个人信息
		u1.GET("/info/get", controller.GetInfo)

		// 修改个人信息
		u1.PUT("/info/update", controller.ModifyInfo)

		// 删除个人信息
		u1.DELETE("/delete", controller.DeleteUser)

		// 修改密码
		u1.PUT("/pwd/update", controller.UpdatePassword)

		// 模糊查询
		u1.GET("/search", controller.FindLike)
	}

	u2 := r.Group("/user")
	{
		// 注册
		u2.POST("/register", controller.Register)
		// 登录
		u2.POST("/login", controller.Login)
	}
}