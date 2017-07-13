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

	beego.Router("/tag/:id([0-9]+)", &controllers.TagController{}, "get:View") //匹配:id 是数字的路由
	beego.Router("/tag/load-more", &controllers.TagController{}, "post:LoadMore")

	//beego.Router("/article", &controllers.ArticlesController{}, "get:Index")
	beego.Router("/article/:id([0-9]+)", &controllers.ArticleController{}, "get:View") //匹配:id 是数字的路由
	beego.Router("/article/load-more", &controllers.ArticleController{}, "post:LoadMore")
	beego.Router("/article/vote", &controllers.ArticleController{}, "post:VoteScore")
	//登陆
	beego.Router("/auth/login", &controllers.AuthController{}, "post:PostLogin")
	beego.Router("/auth/register", &controllers.AuthController{}, "get:GetRegister")
	beego.Router("/auth/register", &controllers.AuthController{}, "post:PostRegister")
	beego.Router("/auth/logout", &controllers.AuthController{}, "get:GetLogout")

	// beego.AutoRouter(&controllers.ImageController{}) //对image使用自动路由
	beego.Router("/subtitle", &controllers.SubtitleController{}, "get:Index")
	beego.Router("/subtitle/load-more", &controllers.SubtitleController{}, "post:LoadMore")

	beego.Router("/video/weibo-parse", &controllers.VideoController{}, "post:WeiboVideoParse")
}
