package model

import (
	"gorm.io/gorm"
)

type Reply struct {
	gorm.Model
	ArticleId uint
	Content   string
}
