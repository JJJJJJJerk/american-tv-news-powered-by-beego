package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "web@gmail.com"
	c.Layout = "layout/base.html"
	c.TplName = "article/index.html"
	c.SetSession("my_name", "awsome")
}
