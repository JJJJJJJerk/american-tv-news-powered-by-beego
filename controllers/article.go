package controllers

import (
	"my_go_web/models"

	"github.com/astaxie/beego"
)

// ArticlesController operations for Articles
type ArticlesController struct {
	beego.Controller
}

// func (c *ArticlesController) GetOne() {
// 	articleID, _ := c.GetInt(":id")
// 	v, err := models.GetArticlesById(articleID)

// 	if err != nil {
// 		//404
// 	} else {
// 		//设置head seo参数
// 		//设置breadcrumb
// 		//设置side bar
// 		//设置head navigation bar
// 		c.Data["Article"] = v
// 		c.Layout = "layout/base.html"
// 		c.TplName = "article/detail.html"
// 	}
// }

func (c *ArticlesController) Index() {
	articles := []models.Article{}
	models.Gorm.Limit(models.PageSize).Order("created_at DESC").Find(&articles)
	c.Data["Articles"] = articles

	c.Data["Title"] = "美剧资讯"
	c.Layout = "layout/base.html"
	c.TplName = "article/index.html"
}
