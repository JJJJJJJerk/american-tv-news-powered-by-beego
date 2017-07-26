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
	WeiboId     uint
	WeiboAvatar string
	WeiboName   string
}

func (user *User) AfterFind() (err error) {
	if user.WeiboAvatar != "" && user.AvatarImage == "" {
		user.AvatarImage = user.WeiboAvatar
	}
	return
}
