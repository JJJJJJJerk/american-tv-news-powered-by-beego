package models

//http://jinzhu.me/gorm/ gorm 文档

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

var CdnHost string

type Image struct {
	gorm.Model
	Key                 string
	Description         string
	ArticleId           uint
	Article             *Article
	Bucket              string
	Fname               string
	Fsize               string
	Width               uint
	Height              uint
	Format              string
	Translation         string
	ImgurId             uint
	Src                 string `gorm:"-"`
	OriginWithWaterMark string `gorm:"-"`
}

func init() {
	CdnHost = beego.AppConfig.String("imageCdnHost")
}

//七牛图片地址转会
func (img *Image) GetImageURL(qiniu string) (url string) {
	url = fmt.Sprintf("%s%s%s", CdnHost, img.Key, qiniu)
	return
}

//quote 图片
func (img *Image) GetQuoteImgURL() (url string) {
	qiniu := "?imageMogr2/gravity/NorthWest/crop/620x350/interlace/1"
	url = fmt.Sprintf("%s%s%s", CdnHost, img.Key, qiniu)
	return
}

func Fetch5RandomQuoteImage() (images []Image) {
	var items []Image
	Gorm.Model(&Image{}).Where("`key` LIKE ?", "%brainyquote%").Order("RAND()").Limit(5).Find(&items)
	return items
}
func Fetch5RandomQuoteImageCached() (images []Image) {

	if x, found := CacheManager.Get(CK_IMG_5_RANDOM); found {
		foo := x.(string)
		buffffer := []byte(foo)
		var items []Image
		json.Unmarshal(buffffer, &items)
		images = items
	} else {
		images = Fetch5RandomQuoteImage()
		data, _ := json.Marshal(images)
		CacheManager.Set(CK_IMG_5_RANDOM, string(data), C_EXPIRE_TIME_HOUR_01)
	}
	return
}

func (image *Image) AfterFind() (err error) {
	//装换excerpt
	qiniu := "?imageMogr2/gravity/NorthWest/crop/620x350/interlace/1"
	image.Src = fmt.Sprintf("%s%s%s", CdnHost, image.Key, qiniu)
	withLogoWaterMark := "?imageView2/0/q/75|watermark/2/text/bW9qb1RW/font/6buR5L2T/fontsize/420/fill/I0ZGRkZGRg==/dissolve/100/gravity/NorthWest/dx/10/dy/10|imageslim"
	image.OriginWithWaterMark = fmt.Sprintf("%s%s%s", CdnHost, image.Key, withLogoWaterMark)
	return
}
