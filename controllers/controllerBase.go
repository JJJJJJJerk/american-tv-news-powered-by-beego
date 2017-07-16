package controllers

import (
	"my_go_web/models"

	"github.com/astaxie/beego"
)

const (
	FlashSuccess    = "flash_success"
	FlashInfo       = "flash_info"
	FlashError      = "flash_Error"
	AuthSessionName = "authed_user_session_name"
)

type BaseController struct {
	beego.Controller //集成beego controller
	//UserInfo         *models.Users
}
type Crumb struct {
	Href  string
	Class string
	Name  string
}

// //为了生成breadcrumb
// type Bread struct {
// 	Name, Href, Class string
// }

func (this *BaseController) Prepare() {

	sessionUser := this.GetSession(AuthSessionName)
	if sessionUser == nil {
		this.Data["User"] = nil
		this.Data["Uid"] = 0
		this.Data["IsAdmin"] = false
	} else {
		user := this.GetSession(AuthSessionName).(models.User)
		this.Data["User"] = user
		this.Data["Uid"] = user.ID
		this.Data["IsAdmin"] = user.ID == 1
	}
	if (this.Ctx.Request.Method == "GET") && !this.Ctx.Input.IsAjax() {
		this.Data["Xsrf"] = this.XSRFToken() //防止跨域
		//fmt.Println(quotes)
		//this.Data["Quotes"] = models.Get3RandomQuote()
		this.Data["Tags"] = models.FetchAllTagsCached()
		this.Data["Imgs"] = models.Fetch5RandomQuoteImageCached()
	}
	// 	//判断用户数是否已近登陆
	// 	//读取session
	// 	userLogin := this.GetSession("loginInfo")
	// 	if userLogin == nil {
	// 		this.Uid = 0
	// 	} else {
	// 		this.UserInfo = userLogin.(*models.Users)
	// 		this.Uid = this.UserInfo.Id
	// 	}
	// 	//在模板里面判断登陆状态
	// 	this.Data["Uid"] = this.Uid
	// 	//做一些权限判断

	//为每一个view 赋值侧边栏

	//fmt.Println(models.Fetch5RandomQuoteImageCached())

	//在这里可以把他填充到模板里面
}

func (this *BaseController) JsonRetrun(status string, message string, data interface{}) {
	this.Data["json"] = map[string]interface{}{"status": status, "message": message, "data": data}
	this.ServeJSON()
	return
}
