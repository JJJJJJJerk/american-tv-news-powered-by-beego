package controllers

import (
	"my_go_web/models"

	"github.com/astaxie/beego"
)

type WxApiController struct {
	beego.Controller //集成beego controller
}

func (c *WxApiController) ArticleIndex() {
	//token = c.GetInt("token")
	offset, _ := c.GetInt(":offset", 0)
	size, _ := c.GetInt(":size", models.PageSize)
	articles := models.GetBatchArticles(offset, size)
	c.Data["json"] = &articles
	c.ServeJSON()
}
func (c *WxApiController) LoadMore() {
	offset, _ := c.GetInt(":offset", 0)
	//tagId, _ := c.GetInt("tagId")
	articles := models.GetBatchArticles(offset, 6)
	c.Data["json"] = &articles
	c.ServeJSON()
}
