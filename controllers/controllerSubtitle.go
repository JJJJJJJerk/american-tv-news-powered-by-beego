package controllers

import (
	"my_go_web/models"
)

type SubtitleController struct {
	BaseController
}

func (c *SubtitleController) Index() {

	var subtitles []models.Subtitle
	models.Gorm.Limit(models.PageSize).Order("created_at DESC").Find(&subtitles, "`uri` <> ''")

	c.Data["BreadCrumbs"] = []Crumb{{"/", "fa fa-home", "首页"}, {"/subtitle/index", "fa fa-home", "资讯"}}
	c.Data["Subtitles"] = subtitles
	c.Data["Keyword"] = "mojoTV字幕,字幕下载,热门美剧,美剧电影,生肉字幕"
	c.Data["Description"] = "mojoTV提供最新欧美美剧资讯,最新预告片,各种趣味小视频,精选制作的美剧字幕"
	c.Data["Title"] = "mojoTV字幕|提供热门美剧资源字幕"

	c.Layout = "layout/base_view.html"
	c.TplName = "subtitle/index.html"
}

func (c *SubtitleController) LoadMore() {
	offset, _ := c.GetInt("offset")
	size, _ := c.GetInt("size")
	var subtitles []models.Subtitle
	models.Gorm.Offset(offset).Limit(size).Order("created_at DESC").Find(&subtitles, "`uri` <> ''")
	c.JsonRetrun("success", "欢迎访问我们的小站", subtitles)
}
