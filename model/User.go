package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model

	Avatar string `gorm:"size:255"`
	Name string `gorm:"type:varchar(20);not null"`
	Email string `gorm:"type:varchar(30);not null;index"`
	Password string `gorm:"size:255;not null"`
	Gender int `gorm:"default:0"`
	Birthday time.Time `gorm:"default:'1970-01-01'"`
	// 个性签名
	Sign string `gorm:"type:varchar(50);default:'这个人很懒，什么都没有留下'"`
	// 1: 管理员，0:普通用户
	Role int `gorm:"1;default:0"`
}