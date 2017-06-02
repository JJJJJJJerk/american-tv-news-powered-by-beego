package main

import (
	//init(register) html-template function ...

	_ "my_go_web/models"
	_ "my_go_web/routers" //init reuters
	"my_go_web/spider"
	_ "my_go_web/tools" //init(register) html-template function ...

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

func main() {
	//spider.RunDygodMeijuSpider()
	//没56分钟27秒执行一次;
	//2两小时一次
	taskSpiderDygodMeiju := toolbox.NewTask("taskSpiderDygodMeiju", "27 * */2 * * *", func() error { spider.RunDygodMeijuSpider(); return nil })
	toolbox.AddTask("taskSpiderDygodMeiju", taskSpiderDygodMeiju)
	toolbox.StartTask()

	beego.Run()

}
