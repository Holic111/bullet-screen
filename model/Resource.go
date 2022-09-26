package model

import "gorm.io/gorm"

// 视频资源表
type Resource struct {
	gorm.Model

	// 所属视频
	Vid uint `gorm:"index"`
	// 分P使用的标题
	Title string `gorm:"type:varchar(50)"`
	// 不同分辨率
	Res360 string `gorm:"type:varchar(255)"`
	Res480 string `gorm:"type:varchar(255)"`
	Res720 string `gorm:"type:varchar(255)"`
	Res1080 string `gorm:"type:varchar(255)"`
	// 不对分辨率进行处理使用原始分辨率
	Original string `gorm:"type:varchar(255)"`
	Duration float64 `gorm:"default:0"`
	Review int `gorm:"not null;index"`
}