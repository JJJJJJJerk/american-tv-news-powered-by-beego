package models

//http://jinzhu.me/gorm/ gorm 文档

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

var CdnHost string

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

func init() {
	CdnHost = beego.AppConfig.String("imageCdnHost")
}

//七牛图片地址转会
func (img *Image) GetImageUrl(qiniu string) (url string) {
	url = fmt.Sprintf("%s%s%s", CdnHost, img.Key, qiniu)
	return
}
