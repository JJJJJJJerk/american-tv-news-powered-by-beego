package controllers

import (
	"fmt"
	"my_go_web/models"
)

// ArticlesController operations for Articles
type ArticlesController struct {
	BaseController
}

func (c *ArticlesController) Index() {

	articles := []models.Article{}
	models.Gorm.Limit(models.PageSize).Order("created_at DESC").Preload("Coverage").Preload("Images").Find(&articles)

	c.Data["BreadCrumbs"] = []Crumb{{"/", "fa fa-home", "首页"}, {"/article", "glyphicon glyphicon-list-alt", "资讯"}}
	c.Data["Articles"] = articles
	c.Data["Keyword"] = "美剧keywords"
	c.Data["Description"] = "美剧描述"
	c.Data["Title"] = "美剧资讯"
	c.Data["Xsrf"] = c.XSRFToken() //防止跨域

	c.Layout = "layout/base.html"
	c.TplName = "article/index.html"
}

func (c *ArticlesController) Detail() {
	articleID, _ := c.GetInt(":id")
	article := models.Article{}
	models.Gorm.First(&article, articleID)

	//设置head seo参数
	//设置breadcrumb
	//设置side bar
	//设置head navigation bar
	url := fmt.Sprintf("/article/%d", articleID)
	c.Data["BreadCrumbs"] = []Crumb{{"/", "fa fa-home", "首页"}, {"/article", "glyphicon glyphicon-list-alt", "资讯"}, {url, "fa fa-graduation-cap", article.Title}}
	c.Data["Article"] = article
	c.Data["Title"] = article.Title

	c.Layout = "layout/base.html"
	c.TplName = "article/detail.html"
}

func (c *ArticlesController) LoadMore() {
	offset, _ := c.GetInt("offset")
	limit := 3
	articles := []models.Article{}
	models.Gorm.Offset(offset).Limit(limit).Order("created_at DESC").Preload("Coverage").Preload("Images").Find(&articles)
	c.JsonRetrun("success", "you are awesome!!!", articles)
}
