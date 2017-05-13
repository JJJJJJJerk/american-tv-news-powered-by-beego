package models

//http://jinzhu.me/gorm/ gorm 文档

import (
	"github.com/jinzhu/gorm"
)

type Image struct {
	gorm.Model
	Key         string
	Description string
	ArticleId   uint
	Article     *Article
	Bucket      string
	Fname       string
	Fsize       string
	Width       uint
	Height      uint
	Format      string
}
