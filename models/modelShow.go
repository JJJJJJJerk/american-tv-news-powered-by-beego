package models

//http://jinzhu.me/gorm/ gorm 文档

import (
	"github.com/jinzhu/gorm"
)
type Show struct {
	gorm.Model
	NameEn string
	NameZh string
}
