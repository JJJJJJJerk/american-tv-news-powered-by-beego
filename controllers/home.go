package controllers

type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {
	c.Layout = "layout/base.html"
	c.TplName = "home/index.html"
}
