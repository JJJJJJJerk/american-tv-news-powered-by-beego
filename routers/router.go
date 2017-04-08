package routers

import (
	"my_go_web/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//注册自动路由
	beego.Router("/article", &controllers.ArticlesController{}, "get:GetAll")
	beego.Router("/article/:id([0-9]+)", &controllers.ArticlesController{}, "get:GetOne") //匹配:id 是数字的路由
	beego.Router("/auth/login", &controllers.AuthController{}, "post:PostLogin")
	beego.Router("/auth/register", &controllers.AuthController{}, "post:PostSignup")

}
