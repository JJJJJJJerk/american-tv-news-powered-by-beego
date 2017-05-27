package models

//http://jinzhu.me/gorm/ gorm 文档

import (
	"github.com/jinzhu/gorm"
)

type ArticleTag struct {
	gorm.Model
	ArticleId uint
	TagId     uint
}

func (ArticleTag) TableName() string {
	return "article_tag"
}
