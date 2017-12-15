package main

import (
	//init(register) html-template function ...

	"www.mojotv.cn/controllers"
	_ "www.mojotv.cn/models"
	_ "www.mojotv.cn/routers"
	_ "www.mojotv.cn/tasks" //启动定时任务
	_ "www.mojotv.cn/tools" //init reuters
	//init(register) html-template function ...

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {

	beego.SetLogger("file", `{"filename":"weibologin.log"}`)

	//spider.RunDygodMovieSpider()
	//没56分钟27秒执行一次;
	//2两小时一次
	//spider.RunDygodMeijuSpider()
	//taskSpiderDygodMeiju := toolbox.NewTask("taskSpiderDygodMeiju", "27 * */2 * * *", func() error { spider.RunDygodMeijuSpider(); return nil })
	//toolbox.AddTask("taskSpiderDygodMeiju", taskSpiderDygodMeiju)
	//toolbox.StartTask()
	beego.ErrorController(&controllers.ErrorController{})
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))
	beego.Run()
}
