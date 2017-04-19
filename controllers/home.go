package controllers

type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {
	//c.Data["BreadCrumb"] = []string
	breadCrumbs := []Bread{{"美剧", "/awesome", "fa fa-dashboard"}, {"how i met your mother", "/how i met", "fa fa-dashboard"}, {"prison", "prison break", "fa fa-dashboard"}}
	c.Data["BreadCrumbs"] = breadCrumbs
	c.Data["Title"] = "首页"
	c.Data["UpdateTime"] = "2017-12-12 12:30:55"
	c.Data["Keyword"] = "美剧,学习英语.权利的游戏,越狱"
	c.Data["Description"] = "最好的美剧学习网站,最好的最好的网站"
	c.Layout = "layout/base.html"
	c.TplName = "home/index.html"
}
