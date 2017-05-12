package models
//http://jinzhu.me/gorm/ gorm 文档

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title    string `gorm:"size:255"`
	Body     string `orm:"column(body)"`
	UrlVideo string
	UrlProvider string
	UrlFlash string
	HtmlCode string
	IsShow uint
	KeyWord string
	MobileCode string
	Discription string
	Images []Image
}
