package models

//http://jinzhu.me/gorm/ gorm 文档

import (
	"github.com/jinzhu/gorm"
)

type ArticleTag struct {
	gorm.Model
	TagId     uint
	ArticleId uint
}

func (User) TableName() string {
	return "article_tag"
}
