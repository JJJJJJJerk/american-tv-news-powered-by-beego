package routers

import (
	"github.com/astaxie/beego"
)

//注册tmeplate的一些实用函数
//https://beego.me/docs/mvc/view/template.md
func init() {
	//注册template function
	beego.AddFuncMap("hi", hello)
}

func hello(in string) (out string) {
	out = in + "world"
	return
}
