package controllers

import (
	"my_go_web/models"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {

	articles := []models.Article{}
	models.Gorm.Limit(models.PageSize).Order("created_at DESC").Preload("Coverage").Preload("Tags").Preload("Images").Find(&articles)

	c.Data["BreadCrumbs"] = []Crumb{{"/", "fa fa-home", "首页"}, {"/article", "fa fa-fire", "资讯"}}
	c.Data["Articles"] = articles
	c.Data["Keyword"] = "美剧keywords"
	c.Data["Title"] = "美剧资讯"

	c.Layout = "layout/base.html"
	c.TplName = "home/index.html"
}
