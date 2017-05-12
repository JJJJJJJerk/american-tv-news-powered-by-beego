package main

import (
	_ "my_go_web/models"  //init(register) html-template function ...
	_ "my_go_web/routers" //init reuters
	_ "my_go_web/tools"   //init(register) html-template function ...

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
