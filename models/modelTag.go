package models

//http://jinzhu.me/gorm/ gorm 文档

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	Name         string
	Description  string
	KeyWord      string
	Image        *Image
	ImageId      uint
	NameEn       string
	Articles     []Article `gorm:"many2many:article_tag;"`
	ArticleCount int       `gorm:"-"`
}

func FetchAllTagsCached() (tags []Tag) {

	if x, found := CacheManager.Get(CK_TAG_ALL); found {
		foo := x.(string)
		buffffer := []byte(foo)
		json.Unmarshal(buffffer, &tags)
	} else {
		now := time.Now().AddDate(0, 0, -7)
		timestring := now.Format("2006-01-02 15:04:05")
		Gorm.Preload("Articles", "articles.created_at >?", timestring).Find(&tags)
		data, _ := json.Marshal(tags)
		CacheManager.Set(CK_TAG_ALL, string(data), C_EXPIRE_TIME_FOREVER)
	}
	return
}

func (tag *Tag) AfterFind() (err error) {
	//装换excerpt
	tag.ArticleCount = len(tag.Articles)
	return
}
