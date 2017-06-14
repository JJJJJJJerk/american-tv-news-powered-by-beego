package models

//http://jinzhu.me/gorm/ gorm 文档

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
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
	if CacheManager.IsExist(CK_QUOTE) {
		value := CacheManager.Get(CK_QUOTE)
		fmt.Println(value)
		data := value.(string)
		fmt.Println(data)
		jsonB := []byte(data)
		json.Unmarshal(jsonB, &quotes)
	} else {
		var count int
		Gorm.Model(&Quote{}).Count(&count)
		rand.Seed(int64(count)) // Try changing this number!
		na := strconv.Itoa(rand.Intn(count))
		nb := strconv.Itoa(rand.Intn(count))
		nc := strconv.Itoa(rand.Intn(count))
		indexs := []string{na, nb, nc}
		Gorm.Model(&Quote{}).Where("id in (?)", indexs).Find(&quotes)
		data, _ := json.Marshal(quotes)
		fmt.Println(string(data))
		CacheManager.Put(CK_QUOTE, string(data), 600*time.Second)

	}
	return

}
