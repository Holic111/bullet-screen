package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model

	Title string `gorm:"type:varchar(50);not null;index"`
	Cover string `gorm:"size:255;not null"`
	Videos []Resource `gorm:"ForeignKey:vid;AssociationForeignKey:id"`
	Desc string `gorm:"type:varchar(200);default:'什么都没有'"` // 视频简介
	Uid uint `gorm:"not null;index"` // 作者ID
	CopyRight bool `gorm:"not null"` // 是否为原创
	Weights float32 `gorm:"default:0"` //视频权重
	Clicks int `gorm:"default:0"` //点击量
	Review int `gorm:"not null"` // 审核状态
	PartitionID uint `gorm:"default:0"` // 分区ID
}