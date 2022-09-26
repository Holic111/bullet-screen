package routers

import (
	"bullet-screen/controller"
	"bullet-screen/middleware"
	"github.com/gin-gonic/gin"
)

type VideoRouter struct {}

func (v *VideoRouter) VideoRouters(r *gin.RouterGroup) {
	v1 := r.Group("/video")
	v1.Use(middleware.JWTAuth())
	{
		// 上传视频信息
		v1.POST("/upload/info", controller.UploadVideoInfo)

		// 修改视频信息
		v1.PUT("/modify/video", controller.UpdateVideoInfo)

		// 获取视频列表
		v1.GET("/upload/get", controller.GetVideoList)

		// 删除视频
		v1.DELETE("/delete", controller.DeleteVideo)

		// 根据分区获取视频列表
		v1.GET("/list/get", controller.GetVideoListByPartitionID)

		// 获取上传资源列表
		v1.GET("/resource/list", controller.GetResourceList)

		// 修改资源标题
		v1.PUT("/resource/title/modify", controller.ModifyResourceTitle)

		// 删除资源
		v1.DELETE("/resource/delete", controller.DeleteResource)
	}
}