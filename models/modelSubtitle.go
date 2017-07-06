package models

//http://jinzhu.me/gorm/ gorm 文档

import (
	"github.com/jinzhu/gorm"
)

type Subtitle struct {
	gorm.Model
	Name string
	Url  string
}
