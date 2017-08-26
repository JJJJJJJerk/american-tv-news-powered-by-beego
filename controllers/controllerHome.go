package controllers

import (
	"my_go_web/models"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {
	articles := models.GetBatchArticles(0, models.PageSize)
	c.Data["BreadCrumbs"] = []Crumb{{"/", "fa fa-home", "首页"}, {"/article", "fa fa-fire", "资讯"}}
	c.Data["Articles"] = articles
	c.Data["Title"] = "mojoTV资讯|最新最快最热的美剧周边资讯"
	c.Data["Keyword"] = "mojoTV资讯,轻松学英语,欧美美剧资讯,国外搞笑小视频,"
	c.Data["Description"] = "mojoTV资讯|提供最新最热最快最热的美剧资讯,海量英语学习资源,欧美搞笑有创意的短视频gif动图,海量美剧双语原生字幕,这里是英语爱好者的乐园,让每一个人都爱上学习英语"
	c.Layout = "layout/base_index.html"
	c.TplName = "home/index.html"
}
