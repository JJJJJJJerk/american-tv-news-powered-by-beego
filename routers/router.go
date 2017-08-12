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
	beego.Router("/toutiao-is-awesome", &controllers.AuthController{}, "get:ToutiaoAd")

	// //注册自动路由

	beego.Router("/tag/:id([0-9]+)", &controllers.TagController{}, "get:View") //匹配:id 是数字的路由
	beego.Router("/tag/load-more", &controllers.TagController{}, "post:LoadMore")
	beego.Router("/tag", &controllers.TagController{}, "post:IndexPost") //获取tag 列表

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

	beego.Router("/rrv/index", &controllers.RenRenController{}, "get:Index")
	beego.Router("/rrv/m3u8/:episodeSid([0-9]+)", &controllers.RenRenController{}, "get:M3u8")
	beego.Router("/rrv/search/:keyword", &controllers.RenRenController{}, "get:Search")
	beego.Router("/rrv/top", &controllers.RenRenController{}, "get:Top")
	beego.Router("/rrv/hot", &controllers.RenRenController{}, "get:Hot")
	beego.Router("/rrv/season/:seasonId([0-9]+)", &controllers.RenRenController{}, "get:Season")
	beego.Router("/rrv/album/:albumId([0-9]+)", &controllers.RenRenController{}, "get:Album")
	beego.Router("/rrv/category/:categoryType/pages/:pages([0-9]+)", &controllers.RenRenController{}, "get:Category")
	//category

	beego.Router("/api/video-parse", &controllers.VideoParseController{}, "get:Index")

}
