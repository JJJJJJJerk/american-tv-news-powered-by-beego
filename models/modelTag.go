package models

//http://jinzhu.me/gorm/ gorm 文档

import (
	"encoding/json"

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

func FetchAllTagsCached() (tags []Tag) {

	if x, found := CacheManager.Get(CK_TAG_ALL); found {
		foo := x.(string)
		buffffer := []byte(foo)
		json.Unmarshal(buffffer, &tags)
	} else {
		Gorm.Find(&tags)
		data, _ := json.Marshal(tags)
		CacheManager.Set(CK_TAG_ALL, string(data), C_EXPIRE_TIME_FOREVER)
	}
	return
}
