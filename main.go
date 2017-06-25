package main

import (
	//init(register) html-template function ...

	_ "my_go_web/models"
	_ "my_go_web/routers" //init reuters
	_ "my_go_web/tools"   //init(register) html-template function ...

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

	beego.Run()

}
