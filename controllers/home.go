package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Layout = "layout/base.html"
	c.TplName = "home/index.html"
}
