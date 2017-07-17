package models

//http://jinzhu.me/gorm/ gorm 文档

import (
	"math"
	"time"

	"strings"

	"regexp"

	"fmt"

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
	Description string
	CoverageUri string
	Images      []Image
	Tags        []Tag  `gorm:"many2many:article_tag;"`
	Shows       []Show `gorm:"many2many:article_show;"`
	Vote        *Vote

	FirstTagName     string `gorm:"_"`
	FirstTagNameEn   string `gorm:"_"`
	FirstTagID       uint   `gorm:"_"`
	Excerpt          string `gorm:"-"`
	CoverageURL      string `gorm:"-"`
	CreatedDate      string `gorm:"-"`
	CreatedTime      string `gorm:"-"`
	CreatedHumanTime string `gorm:"-"`
	VideoYoukuId     string `gorm:"-"`
	VideoMiaopaiId   string `gorm:"-"`
	VideoWeiboId     string `gorm:"-"`
	Links            []Link `gorm:"-"`
}

const ShowHost = "//www.trytv.org/"

//做一些计算
func (art *Article) AfterFind() (err error) {
	//装换excerpt
	body := beego.HTML2str(art.Body)
	art.Excerpt = beego.Substr(body, 0, 120)
	//转换时间啊
	art.CreatedDate = beego.Date(art.CreatedAt, "m-d")
	art.CreatedTime = beego.Date(art.CreatedAt, "H:i")
	art.CreatedHumanTime = CovertTimeToHumanTime(art.CreatedAt)
	//param := "?imageView2/1/w/120/h/120"
	//param := "?imageView2/1/w/480/h/270"
	param := "?imageMogr2/thumbnail/!620x350r/gravity/Center/crop/480x270/blur/1x0/quality/80|watermark/2/text/bW9qb3R2/font/6buR5L2T/fontsize/360/fill/I0ZGRkZGRg==/dissolve/100/gravity/SouthWest/dx/10/dy/10|imageslim"

	imageModel := Image{Key: "article-placeholder"}
	if art.CoverageUri != "" {
		imageModel.Key = art.CoverageUri
	}
	art.CoverageURL = imageModel.GetImageURL(param)

	//解析视频类型的优酷vid
	if strings.Contains(art.UrlVideo, "youku.com") {
		//http://v.youku.com/v_show/id_XMjg4Mzc0NjAxMg==.html?spm=a2hww.20023042.m_223465.5~5~5~5!2~5~5~A&f=50346975
		//http://m.youku.com/video/id_XMjg4Mzc0NjAxMg==.html?spm=a2hww.20023042.m_223465.5~5~5~5!2~5~5~A&f=50346975&source=
		reg := regexp.MustCompile(`(?:id_)(\w+)(?:[=]{0,2}\.html)`)
		ids := reg.FindStringSubmatch(art.UrlVideo)
		for k, v := range ids {
			if k == 1 {
				art.VideoYoukuId = v
			}
		}
	}
	//解析微博
	//https://m.weibo.cn/status/Fc99eEAbb?fid=1034%3Ae4cb370b2f219a79e8e0d55a4a3bb673&jumpfrom=weibocom
	//http://weibo.com/tv/v/Fc99eEAbb?fid=1034:e4cb370b2f219a79e8e0d55a4a3bb673
	//http://weibo.com/tv/v/Fc13IAmqT?fid=1034:94f6f34920fa8d2a353f85b6f3fde66a
	//http://weibo.com/tv/v/Fc2Mudq0n?fid=1034:2534bf1ffad6701f1a4ac7e60575c756
	//http://weibo.com/tv/v/Fc4DPrLRJ?from=vhot
	//http://weibo.com/tv/v/Fc99eEAbb?fid=1034:e4cb370b2f219a79e8e0d55a4a3bb673

	if strings.Contains(art.UrlVideo, "weibo.") {
		//http://v.youku.com/v_show/id_XMjg4Mzc0NjAxMg==.html?spm=a2hww.20023042.m_223465.5~5~5~5!2~5~5~A&f=50346975
		//http://m.youku.com/video/id_XMjg4Mzc0NjAxMg==.html?spm=a2hww.20023042.m_223465.5~5~5~5!2~5~5~A&f=50346975&source=
		reg := regexp.MustCompile(`\?fid=(\d{4}:\w{32})`)
		ids := reg.FindStringSubmatch(art.UrlVideo)
		for k, v := range ids {
			if k == 1 {
				art.VideoWeiboId = v
			}
		}
	}
	//解析秒拍
	//http://www.miaopai.com/show/guASDNtbED2~Q-G9lBSCx1ECxxj~vqCc.htm
	//http://www.miaopai.com/show/XvEqhME836J9MvRgh6xFz~BY5PNjYS~e.htm
	//http://gslb.miaopai.com/stream/c48Xw9OZhOmLRaQeG1vqIxFax5Pmp29O.mp4?yx=&refer=weibo_app&Expires=1499883561&ssig=ibesKAxQRj&KID=unistore,video&playerType=miaopai
	if strings.Contains(art.UrlVideo, "miaopai.") {
		reg := regexp.MustCompile(`(?:/show/)(.+)(?:\.htm)`)
		ids := reg.FindStringSubmatch(art.UrlVideo)
		for k, v := range ids {
			if k == 1 {
				art.VideoMiaopaiId = v
			}
		}
	}

	//处理标签
	for k, v := range art.Tags {
		//设置第一个标签
		if k == 0 {
			art.FirstTagID = v.ID
			art.FirstTagName = v.Name
			art.FirstTagNameEn = v.NameEn
		}
		//设置外链内链
		url := fmt.Sprint("/tag/", v.ID)
		link := Link{Name: v.Name, Url: url}
		art.Links = append(art.Links, link)
	}
	for _, show := range art.Shows {
		url := fmt.Sprint(ShowHost, "show/", show.ID)
		showName := fmt.Sprint(show.NameEn, show.NameZh)
		link := Link{Name: showName, Url: url}
		art.Links = append(art.Links, link)
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

func CovertTimeToHumanTime(t time.Time) (humanTimeString string) {
	secT := t.Unix()
	secNow := time.Now().Unix()
	duration := secNow - secT
	if duration < 60 {
		humanTimeString = "现在"
		return
	} else if duration < 3600 {
		humanTimeString = fmt.Sprint(duration/60, "分前")
		return
	} else if duration < 24*3600 {
		humanTimeString = fmt.Sprint(duration/3600, "小时前")
		return
	} else if duration < 24*3600*30 {
		humanTimeString = fmt.Sprint(duration/24/3600, "天前")
		return
	} else if duration < 24*3600*366 {
		humanTimeString = fmt.Sprint(duration/24/3600/30, "月前")
		return
	} else {
		humanTimeString = t.Format("2006-01-02")
		return
	}
}
