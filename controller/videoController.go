package controller

import (
	"bullet-screen/model"
	"bullet-screen/model/public"
	"bullet-screen/model/relative"
	"bullet-screen/service"
	"bullet-screen/util"
	"github.com/gin-gonic/gin"
)

// 上传视频信息
func UploadVideoInfo(c *gin.Context) {
	var v relative.Info_Video
	err := c.ShouldBindJSON(&v)
	if err != nil {
		return
	}

	var video = model.Video{
		Title:       v.Title,
		Cover:       v.Cover,
		Desc:        v.Desc,
		CopyRight:   v.Original,
		PartitionID: uint(v.Partition),
	}

	vid, code := service.UploadVideoInfo(video)

	util.ResponseData(code, vid, c)
}

// 修改视频信息
func UpdateVideoInfo(c *gin.Context) {
	var v relative.Update_Video
	err := c.ShouldBindJSON(&v)
	if err != nil {
		return
	}

	video := model.Video{
		Title: v.Title,
		Cover: v.Cover,
		Desc: v.Desc,
		CopyRight: v.Original,
	}

	video.ID = uint(v.Vid)

	code := service.UpdateVideoInfo(video)

	util.ResponseData(code, nil, c)
}

// 获取上传视频列表
func GetVideoList(c *gin.Context) {
	var page public.Page
	page.PageNum, _ = util.StringToInt(c.Query("page"))
	page.PageSize, _ = util.StringToInt(c.Query("page_size"))

	videos, code := service.GetVideoList(page)
	if code != util.OK {
		util.ResponseData(code, nil, c)
		return
	}
	util.ResponseData(code, videos, c)
}

// 通过分区获取所有视频列表，分区id为0时，获取所有分区视频
func GetVideoListByPartitionID(c *gin.Context) {
	id, err := util.StringToInt(c.Query("partition"))
	if err != nil {
		return
	}

	videos, code := service.GetVideoListByPartitionID(id)

	if code != util.OK {
		util.ResponseData(code, nil, c)
		return
	}

	util.ResponseData(code, videos, c)
}


// 删除视频
func DeleteVideo(c *gin.Context) {
	id, _ := util.StringToInt(c.PostForm("id"))

	code := service.DeleteVideoById(id)

	util.ResponseData(code, nil, c)
}

// 获取上传资源列表
func GetResourceList(c *gin.Context) {
	var vid int
	vid, err := util.StringToInt(c.Query("vid"))
	if err != nil {
		return
	}

	resource, code := service.GetResourceList(vid)
	if code != util.OK {
		util.ResponseData(code, nil, c)
		return
	}
	util.ResponseData(code, resource, c)
}

// 修改资源标题
func ModifyResourceTitle(c *gin.Context) {
	var r relative.Modify_Resource
	err := c.ShouldBindJSON(&r)
	if err != nil {
		return
	}
	code := service.ModifyResourceTitle(r)

	util.ResponseData(code, nil, c)
}

// 删除资源
func DeleteResource(c *gin.Context) {
	id, err := util.StringToInt(c.Query("id"))
	if err != nil {
		return
	}

	code := service.DeleteResource(id)

	util.ResponseData(code, nil, c)
}