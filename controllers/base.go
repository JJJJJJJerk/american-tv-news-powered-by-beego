package controllers

import (
	"github.com/astaxie/beego"
)

const (
	FlashSuccess = "flash_success"
	FlashInfo    = "flash_info"
	FlashError   = "flash_Error"
)

type BaseController struct {
	beego.Controller //集成beego controller
	Uid              int
	//UserInfo         *models.Users
}

// //为了生成breadcrumb
// type Bread struct {
// 	Name, Href, Class string
// }

// func (this *BaseController) Prepare() {
// 	this.Data["app_name"] = beego.AppConfig.String("appname")
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
// }
