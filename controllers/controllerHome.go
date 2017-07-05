package controllers

import (
	"my_go_web/models"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {
	articles := []models.Article{}
	models.Gorm.Limit(models.PageSize).Order("created_at DESC").Preload("Coverage").Preload("Vote").Preload("Tags").Preload("Images").Find(&articles)
	c.Data["BreadCrumbs"] = []Crumb{{"/", "fa fa-home", "首页"}, {"/article", "fa fa-fire", "资讯"}}
	c.Data["Articles"] = articles
	c.Data["Keyword"] = "提供最新最热的美剧资讯,美剧播放时间表,欧美搞笑有创意的短视频gif动图,英文电子有声读物,海量美剧双语原生字幕,这里是英语爱好者的乐园,让每一个人都爱上学习英语"
	c.Data["Title"] = "娱乐学习英语"
	c.Data["Keyword"] = "轻松学英语,欧美,搞笑gif动图短片,"
	c.Data["Description"] = ""
	c.Layout = "layout/base_index.html"
	c.TplName = "home/index.html"
}
