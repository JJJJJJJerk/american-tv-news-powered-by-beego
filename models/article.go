package models

//http://jinzhu.me/gorm/ gorm 文档

import (
	"math"

	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title       string `gorm:"size:255"`
	Body        string `orm:"column(body)"`
	UrlVideo    string
	UrlProvider string
	UrlFlash    string
	HtmlCode    string
	IsShow      uint
	KeyWord     string
	MobileCode  string
	Discription string
	CoverageId  uint32
	Images      []Image
}

func GetAllArticles(pageIndex int) (articles []Article, totalPage int) {
	//设置默认值
	if pageIndex < 1 {
		pageIndex = 1
	}

	//分页
	var count int
	Gorm.Model(&Article{}).Count(&count)
	totalPage = int(math.Ceil(float64(count) / float64(PageSize)))

	offset := (pageIndex - 1) * PageSize
	articles = []Article{}
	Gorm.Offset(offset).Limit(PageSize).Order("created_time DESC").Find(&articles)
	return
}
