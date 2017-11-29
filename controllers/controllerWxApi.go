package controllers

import (
	"www.mojotv.cn/models"

	"github.com/astaxie/beego"
)

type WxApiController struct {
	beego.Controller //集成beego controller
}

func (c *WxApiController) ArticleIndex() {
	//token = c.GetInt("token")
	offset, _ := c.GetInt(":offset", 0)
	size, _ := c.GetInt(":size", models.PageSize)
	articles := models.GetBatchArticlesForWx(offset, size)
	c.Data["json"] = &articles
	c.ServeJSON()
}
