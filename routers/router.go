package routers

import (
	"my_go_web/controllers"

	"github.com/astaxie/beego"
)

//功能一:使用电视剧角色和名称图像
//功能二:发帖功能
//功能三:提供小视屏播放功能
//仿照斗鱼打赏功能
//
func init() {
	beego.Router("/", &controllers.HomeController{}, "get:Get")
	// //注册自动路由
	//beego.Router("/article", &controllers.ArticlesController{}, "get:Index")
	beego.Router("/article/:id([0-9]+)", &controllers.ArticlesController{}, "get:Detail") //匹配:id 是数字的路由
	beego.Router("/article/load-more", &controllers.ArticlesController{}, "post:LoadMore")
	beego.Router("/article/vote", &controllers.ArticlesController{}, "post:VoteScore")
	//登陆
	beego.Router("/auth/login", &controllers.AuthController{}, "post:PostLogin")
	beego.Router("/auth/register", &controllers.AuthController{}, "get:GetRegister")
	beego.Router("/auth/register", &controllers.AuthController{}, "post:PostRegister")

	// beego.AutoRouter(&controllers.ImageController{}) //对image使用自动路由

}
