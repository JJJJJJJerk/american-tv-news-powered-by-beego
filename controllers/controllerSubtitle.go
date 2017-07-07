package controllers

import (
	"my_go_web/models"
)

type SubtitleController struct {
	BaseController
}

func (c *SubtitleController) Index() {
	c.Data["Tags"] = models.FetchAllTagsCached()

	var subtitles []models.Subtitle
	models.Gorm.Limit(models.PageSize).Order("created_at DESC").Find(&subtitles, "`uri` <> ''")

	c.Data["BreadCrumbs"] = []Crumb{{"/", "fa fa-home", "首页"}, {"/subtitle/index", "fa fa-home", "资讯"}}
	c.Data["Subtitles"] = subtitles
	c.Data["Keyword"] = "美剧keywords"
	c.Data["Description"] = "美剧描述"
	c.Data["Title"] = "美剧资讯"

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
