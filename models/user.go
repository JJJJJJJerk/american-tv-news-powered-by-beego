package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name        string `gorm:"size:255"`
	Email       string `gorm:"size:255"`
	Password    string
	AvatarImage string
}
