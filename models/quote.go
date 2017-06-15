package models

//http://jinzhu.me/gorm/ gorm 文档

import (
	"encoding/json"
	"math/rand"
	"strconv"

	"github.com/jinzhu/gorm"
	cache "github.com/patrickmn/go-cache"
)

const CK_QUOTE = "CacheKey.3randomQuotes"

type Quote struct {
	gorm.Model
	English  string `json:"english";`
	Chinese  string `json:"chinese";`
	Writer   string `gorm:"size:255";json:"writer";`
	ImageUri string `orm:"column(body)";json:"image_uri";`
}

func Get3RandomQuote() (quotes []Quote) {
	if x, found := CacheManager.Get(CK_QUOTE); found {
		foo := x.(string)
		buffffer := []byte(foo)
		var items []Quote
		json.Unmarshal(buffffer, &items)
		quotes = items
	} else {
		quotes = QuoteRandom3()
		data, _ := json.Marshal(quotes)
		CacheManager.Set(CK_QUOTE, string(data), cache.DefaultExpiration)
	}
	return
}

func QuoteRandom3() (quotes []Quote) {
	var count int
	Gorm.Model(&Quote{}).Count(&count)
	rand.Seed(int64(count)) // Try changing this number!
	na := strconv.Itoa(rand.Intn(count))
	nb := strconv.Itoa(rand.Intn(count))
	nc := strconv.Itoa(rand.Intn(count))
	nd := strconv.Itoa(rand.Intn(count))
	ne := strconv.Itoa(rand.Intn(count))
	indexs := []string{na, nb, nc, nd, ne}
	Gorm.Model(&Quote{}).Where("id in (?)", indexs).Find(&quotes)
	return
}
