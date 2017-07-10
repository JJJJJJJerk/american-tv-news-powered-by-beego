package controllers

import (
	"github.com/astaxie/beego"
)

// ArticlesController operations for Articles
type ImageController struct {
	beego.Controller
}

// func (c *ImageController) GetOne() {
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

// func (c *ImageController) Upload() {
// 	_, articles := models.GetAllArticles()
// 	c.Data["Articles"] = articles
// 	c.Data["Title"] = "aweosme go web application!!!!"
// 	c.Layout = "layout/base.html"
// 	c.TplName = "article/index.html"
// }

// func (c *ImageController) Upload_canvas() {
// 	_, articles := models.GetAllArticles()
// 	c.Data["Articles"] = articles
// 	c.Data["Title"] = "aweosme go web application!!!!"
// 	c.Layout = "layout/base.html"
// 	c.TplName = "image/cropper_modal.html"
// }
