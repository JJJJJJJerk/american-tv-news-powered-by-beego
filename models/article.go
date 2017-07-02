package models

//http://jinzhu.me/gorm/ gorm 文档

import (
	"math"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	RawTitle    string
	RawContent  string
	Title       string `gorm:"size:255"`
	Body        string `orm:"column(body)"`
	UrlVideo    string
	UrlProvider string
	VideoCode   string
	IsShow      uint
	KeyWord     string
	Discription string
	CoverageId  uint32
	Coverage    *Image `gorm:"ForeignKey:CoverageId"`
	Images      []Image
	ReadCount   uint16

	Excerpt        string `gorm:"-"` //计算出文章摘要
	CoverageUrl    string `gorm:"-"` //文章封面
	CreatedDate    string `gorm:"-"`
	CreatedTime    string `gorm:"-"`
	Tags           []Tag  `gorm:"many2many:article_tag;"`
	Vote           Vote
	FirstTagName   string `gorm:"_";` //第一个标签的名称
	FirstTagNameEn string `gorm:"_";` //第一个标签的名称
	FirstTagID     uint   `gorm:"_";` //第一个标签的名称
}

//做一些计算
func (art *Article) AfterFind() (err error) {
	//装换excerpt
	body := beego.HTML2str(art.Body)
	art.Excerpt = beego.Substr(body, 0, 120)
	//转换时间啊
	art.CreatedDate = beego.Date(art.CreatedAt, "m-d")
	art.CreatedTime = beego.Date(art.CreatedAt, "H:i")

	//param := "?imageView2/1/w/120/h/120"
	param := "?imageView2/1/w/480/h/270"

	if art.Coverage != nil {
		art.CoverageUrl = art.Coverage.GetImageUrl(param)
		return

	}
	if len(art.Images) > 0 {
		art.CoverageUrl = art.Images[0].GetImageUrl(param)
		return

	}
	defaultImage := Image{Key: "1461329417"}
	art.CoverageUrl = defaultImage.GetImageUrl(param)

	art.FirstTagName = art.Tags[0].Name
	art.FirstTagNameEn = art.Tags[0].NameEn
	art.FirstTagID = art.Tags[0].ID
	return
}

//这个方法要被废弃了
func GetAllArticles(pageIndex int) (articles []Article, totalPage int) {
	//设置默认值
	if pageIndex < 1 {
		pageIndex = 1
	}

	//分页
	var count int
	Gorm.Model(&Article{}).Count(&count)
	totalPage = int(math.Ceil(float64(count) / float64(PageSize)))

	offset := (pageIndex - 1) * PageSize
	articles = []Article{}
	Gorm.Offset(offset).Limit(PageSize).Order("created_time DESC").Preload("Vote").Find(&articles)
	return
}
