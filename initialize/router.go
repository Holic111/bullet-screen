package initialize

import (
	"bullet-screen/common"
	"bullet-screen/middleware"
	"bullet-screen/routers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RouterInit() {
	gin.SetMode(common.Global_AppMode.Mode)
	engine := gin.New()

	entryRouter := new(routers.EnterRouters)

	//engine.Static("/img", "./img/")

	v1 := engine.Group("api/v1")

	v1.Use(middleware.ZapLogger())
	v1.Use(middleware.ZapRecovery(true))

	{
		entryRouter.UserRouters(v1)

		entryRouter.CaptchaRouters(v1)

		entryRouter.VideoRouters(v1)
	}




	err := engine.Run(common.Global_AppMode.Port)
	if err != nil {
		fmt.Println(err)
	}
}