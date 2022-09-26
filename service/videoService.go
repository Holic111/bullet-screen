package service

import (
	"bullet-screen/common"
	"bullet-screen/model"
	"bullet-screen/model/public"
	"bullet-screen/model/relative"
	"bullet-screen/util"
)

func UploadVideoInfo(video model.Video) (int, int) {
	err := common.Global_Mysql.Debug().Create(&video).Error
	if err != nil {
		return -1, util.ERROR
	}
	return int(video.ID), util.OK
}

func UpdateVideoInfo(video model.Video) int {
	err := common.Global_Mysql.Debug().Table("video").Where("id = ?", video.ID).Updates(&video).Error
	if err != nil {
		return util.ERROR
	}
	return util.OK
}

func GetVideoList(page public.Page) ([]*model.Video, int) {
	if page.PageNum <= 0 {
		page.PageNum = 1
	}
	if page.PageSize <=0 {
		page.PageSize = 10
	}
	var videos []*model.Video
	err := common.Global_Mysql.Debug().Table("video").Limit(page.PageSize).Offset(page.PageSize * (page.PageNum-1)).Find(&videos).Error
	if err != nil {
		return nil, util.ERROR
	}
	return videos, util.OK
}

func DeleteVideoById(id int) int {
	var u model.Video
	err := common.Global_Mysql.Debug().Where("id = ?", id).Delete(&u).Error
	if err != nil {
		return util.ERROR
	}
	return util.OK
}

func GetResourceList(vid int) ([]*model.Resource, int) {
	var resource []*model.Resource
	err := common.Global_Mysql.Debug().Where("vid = ?", vid).Find(&resource).Error
	if err != nil {
		return nil, util.ERROR
	}
	return resource, util.OK
}

func ModifyResourceTitle(resource relative.Modify_Resource) int {
	err := common.Global_Mysql.Debug().Table("resource").Where("id = ?", resource.ID).Update("title", resource.Title).Error
	if err != nil {
		return util.ERROR
	}
	return util.OK
}

func GetVideoListByPartitionID(id int) ([]*relative.Paratition_Video, int) {
	var video []*relative.Paratition_Video
	err := common.Global_Mysql.Debug().Table("video").Where("partition_id = ?", id).Find(&video).Error
	if err != nil {
		return nil, util.ERROR
	}
	return video, util.OK
}

func DeleteResource(id int) int {
	err := common.Global_Mysql.Debug().Table("resource").Where("id = ?", id).Delete(&model.Resource{}).Error
	if err != nil {
		return util.ERROR
	}
	return util.OK
}
