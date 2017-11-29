package models

import (
	"github.com/jinzhu/gorm"
)

type Episode struct {
	gorm.Model
	Name      string `gorm:"size:255"`
	RawName   string `gorm:"size:255"`
	Provider  string
	UrlMagnet string
}
