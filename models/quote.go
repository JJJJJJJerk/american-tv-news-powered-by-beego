package models

//http://jinzhu.me/gorm/ gorm 文档

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

const CK_QUOTE = "CacheKey.3randomQuotes"

type Quote struct {
	gorm.Model
	English  string
	Chinese  string
	Writer   string `gorm:"size:255"`
	ImageUri string `orm:"column(body)"`
}

func Get3RandomQuote() (quotes []Quote) {
	if CacheManager.IsExist(CK_QUOTE) {
		value := CacheManager.Get(CK_QUOTE)
		fmt.Println(value)
		quotes = CacheManager.Get(CK_QUOTE).([]Quote)
	} else {
		var count int
		Gorm.Model(&Quote{}).Count(&count)
		rand.Seed(int64(count)) // Try changing this number!
		na := strconv.Itoa(rand.Intn(count))
		nb := strconv.Itoa(rand.Intn(count))
		nc := strconv.Itoa(rand.Intn(count))
		indexs := []string{na, nb, nc}
		Gorm.Model(&Quote{}).Where("id in (?)", indexs).Find(&quotes)
		CacheManager.Put(CK_QUOTE, quotes, 600*time.Second)

	}
	return

}
