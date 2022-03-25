package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name string
}

type TagPost struct {
	PostId uint
	TagId  uint
}
