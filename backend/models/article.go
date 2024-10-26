package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string
	Content string
	Preview string
	Likes   int `gorm:"default:0"` // 点赞数，默认为0
}
