package main

import (
	//init(register) html-template function ...

	"my_go_web/controllers"
	_ "my_go_web/models"
	_ "my_go_web/routers"
	_ "my_go_web/tools" //init reuters
	//init(register) html-template function ...

	"github.com/astaxie/beego"
)

func main() {

	beego.SetLogger("file", `{"filename":"test.log"}`)

	//spider.RunDygodMovieSpider()
	//没56分钟27秒执行一次;
	//2两小时一次
	//spider.RunDygodMeijuSpider()
	//taskSpiderDygodMeiju := toolbox.NewTask("taskSpiderDygodMeiju", "27 * */2 * * *", func() error { spider.RunDygodMeijuSpider(); return nil })
	//toolbox.AddTask("taskSpiderDygodMeiju", taskSpiderDygodMeiju)
	//toolbox.StartTask()
	beego.ErrorController(&controllers.ErrorController{})

	beego.Run()

}
