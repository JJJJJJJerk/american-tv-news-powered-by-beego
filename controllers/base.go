package controllers

import (
	"my_go_web/models"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller //集成beego controller
	Uid              int
	UserInfo         *models.Users
}

func (this *BaseController) Prepare() {
	//判断用户数是否已近登陆
	//读取session
	userLogin := this.GetSession("loginInfo")
	if userLogin == nil {
		this.Uid = 0
	} else {
		this.UserInfo = userLogin.(*models.Users)
		this.Uid = this.UserInfo.Id
	}
	//做一些权限判断
}
