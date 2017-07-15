package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["content"] = "page not found"
	c.TplName = "error/404.html"
}

func (c *ErrorController) Error501() {
	c.Data["content"] = "server error"
	c.TplName = "error/404.html"
}
func (c *ErrorController) Error503() {
	c.Data["content"] = "server error"
	c.TplName = "error/404.html"
}
func (c *ErrorController) Error500() {
	c.Data["content"] = "server error"
	c.TplName = "error/404.html"
}
func (c *ErrorController) Error401() {
	c.Data["content"] = "server error"
	c.TplName = "error/404.html"
}

func (c *ErrorController) Error403() {
	c.Data["content"] = "server error"
	c.TplName = "error/404.html"
}
func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "error/404.html"
}
