package models

//http://jinzhu.me/gorm/ gorm 文档

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	Name        string
	Description string
	KeyWord     string
	Image       *Image
	ImageId     uint
	NameEn      string
	Articles    []Article `gorm:"many2many:article_tag;"`
}
