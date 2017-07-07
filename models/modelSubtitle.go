package models

//http://jinzhu.me/gorm/ gorm 文档

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Subtitle struct {
	gorm.Model
	NameZh    string
	NameEn    string
	Version   string
	Format    string
	SourceUrl string
	FileName  string
	Lang      string
	Uri       string
	Url       string
	OssUrl    string `gorm:"-"`
	HumamTime string `gorm:"_"`
}

func (sub *Subtitle) AfterFind()(err error) {
	sub.OssUrl = fmt.Sprintf("%s%s", CdnHost, sub.Uri)
	sub.HumamTime = sub.CreatedAt.Format("06-01-02 15:04")
	return
}
