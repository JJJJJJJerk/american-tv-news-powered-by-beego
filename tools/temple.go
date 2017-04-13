package routers

//防止模板相关的方便函数
import (
	"fmt"

	"github.com/astaxie/beego"
)

//注册tmeplate的一些实用函数
//https://beego.me/docs/mvc/view/template.md

//配置cd镜像地址
var cdn string

func init() {
	cdn = beego.AppConfig.DefaultString("cdnhost", "/static/")

	//注册template function
	beego.AddFuncMap("hi", hello)
	beego.AddFuncMap("cdnSrc", cdnSrc) //拼接cdn资源
}

func hello(in string) (out string) {
	out = in + "world"
	return
}

//得到cdn的完全地址
//cdn=http://static.trytv.org/
//src=app.css
//return http://static.trytv.org/app.css
func cdnSrc(in string) (out string) {
	out = fmt.Sprintf("%s%s", cdn, in)
	return
}
