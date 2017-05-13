package models

//http://jinzhu.me/gorm/ gorm 文档

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

var imageCdnHost string

func init() {
	imageCdnHost = beego.AppConfig.String("imageCdnHost")
	fmt.Println(imageCdnHost)
}

type Image struct {
	gorm.Model
	Key         string
	Description string
	ArticleId   uint
	Article     Article
	Bucket      string
	Fname       string
	Fsize       string
	Width       uint
	Height      uint
	Format      string
}

func (image *Image) Url(qiniuParameter string) string {
	return fmt.Sprintf("%s%s%s", imageCdnHost, image.Key, qiniuParameter)
}
