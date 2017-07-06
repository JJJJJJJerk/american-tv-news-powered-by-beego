package controllers

import (
	"fmt"
	"my_go_web/models"
)

type TagController struct {
	BaseController
}

func (c *TagController) Index() {

	articles := []models.Article{}
	models.Gorm.Limit(models.PageSize).Order("created_at DESC").Preload("Vote").Preload("Tags").Preload("Images").Find(&articles)

	c.Data["BreadCrumbs"] = []Crumb{{"/", "fa fa-home", "首页"}, {"/article", "fa fa-home", "资讯"}}
	c.Data["Articles"] = articles
	c.Data["Keyword"] = "美剧keywords"
	c.Data["Description"] = "美剧描述"
	c.Data["Title"] = "美剧资讯"

	c.Layout = "layout/base.html"
	c.TplName = "article/index.html"
}

func (c *TagController) View() {
	tagId, _ := c.GetInt(":id")
	//浏览计数
	tag := models.Tag{}
	var articles []models.Article
	models.Gorm.First(&tag, tagId)

	//models.Gorm.Related("Tags", "article_tag.tag_id = ?", tag.ID).Preload("Images").Limit(90).Find(&articles)
	models.Gorm.Model(&tag).Order("articles.updated_at desc").Limit(models.PageSize).Preload("Images").Preload("Vote").Related(&articles, "Articles")

	//设置head seo参数
	//设置breadcrumb
	//设置side bar
	//设置head navigation bar
	url := fmt.Sprintf("/tag/%d", tagId)
	c.Data["BreadCrumbs"] = []Crumb{{"/", "fa fa-home", "首页"}, {url, "fa fa-navicon", tag.Name}}
	c.Data["Tag"] = tag
	c.Data["Articles"] = articles
	c.Data["Title"] = tag.Name
	c.Data["Keyword"] = tag.KeyWord
	c.Data["Description"] = tag.Description

	c.Layout = "layout/base_index.html"
	c.TplName = "tag/view.html"
}

func (c *TagController) LoadMore() {
	offset, _ := c.GetInt("offset")
	size, _ := c.GetInt("size")
	tagId, _ := c.GetInt("tagId")
	tag := models.Tag{}
	tag.ID = uint(tagId)
	articles := []models.Article{}
	models.Gorm.Model(&tag).Offset(offset).Limit(size).Order("articles.updated_at DESC").Preload("Images").Preload("Tags").Preload("Vote").Related(&articles, "Articles")
	c.JsonRetrun("success", "欢迎访问我们的小站", articles)
}
