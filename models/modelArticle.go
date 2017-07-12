package models

//http://jinzhu.me/gorm/ gorm 文档

import (
	"math"

	"strings"

	"regexp"

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
	CoverageUri string
	Images      []Image
	Tags        []Tag `gorm:"many2many:article_tag;"`
	Vote        *Vote

	FirstTagName   string `gorm:"_"`
	FirstTagNameEn string `gorm:"_"`
	FirstTagID     uint   `gorm:"_"`
	Excerpt        string `gorm:"-"`
	CoverageURL    string `gorm:"-"`
	CreatedDate    string `gorm:"-"`
	CreatedTime    string `gorm:"-"`
	CreatedWeekday string `gorm:"-"`
	VideoYoukuId   string `gorm:"-"`
}

//做一些计算
func (art *Article) AfterFind() (err error) {
	//装换excerpt
	body := beego.HTML2str(art.Body)
	art.Excerpt = beego.Substr(body, 0, 120)
	//转换时间啊
	art.CreatedDate = beego.Date(art.CreatedAt, "m-d")
	art.CreatedTime = beego.Date(art.CreatedAt, "H:i")
	art.CreatedWeekday = art.CreatedAt.Format("Mon 15:04")
	//param := "?imageView2/1/w/120/h/120"
	//param := "?imageView2/1/w/480/h/270"
	param := "?imageMogr2/auto-orient/thumbnail/!480x270r/gravity/NorthWest/crop/480x270/format/png/blur/1x0/quality/100|imageslim"
	if len(art.Tags) > 0 {
		firstTag := art.Tags[0]
		art.FirstTagID = firstTag.ID
		art.FirstTagName = firstTag.Name
		art.FirstTagNameEn = firstTag.NameEn
	}
	imageModel := Image{Key: "article-placeholder"}
	if art.CoverageUri != "" {
		imageModel.Key = art.CoverageUri
	}
	art.CoverageURL = imageModel.GetImageURL(param)

	//解析视频类型的vid
	if strings.Contains(art.UrlVideo, "youku.com") {
		//http://v.youku.com/v_show/id_XMjg4Mzc0NjAxMg==.html?spm=a2hww.20023042.m_223465.5~5~5~5!2~5~5~A&f=50346975
		//http://m.youku.com/video/id_XMjg4Mzc0NjAxMg==.html?spm=a2hww.20023042.m_223465.5~5~5~5!2~5~5~A&f=50346975&source=
		reg := regexp.MustCompile(`(?:id_)(\w+)(?:==\.html)`)
		ids := reg.FindStringSubmatch(art.UrlVideo)
		for k, v := range ids {
			if k == 1 {
				art.VideoYoukuId = v
			}
		}
	}

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
