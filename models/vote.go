package models

//http://jinzhu.me/gorm/ gorm 文档

import (
	"github.com/jinzhu/gorm"
)

type Vote struct {
	gorm.Model
	ArticleId     uint16
	ShowId        uint16
	MovieId       uint16
	Visit         uint16
	Score         uint16
	VoteCount     uint16
	FavorateCount uint16
}
