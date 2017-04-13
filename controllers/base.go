package controllers

import (
	"my_go_web/models"

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
	//在模板里面判断登陆状态
	this.Data["Uid"] = this.Uid
	//做一些权限判断

	//处理flash session

	this.Data[FlashError] = this.GetSession(FlashError).([]string)
	this.DelSession(FlashError)
}

func (this *BaseController) FlashError(messages []string) {
	this.SetSession(FlashError, messages)
}
func (this *BaseController) FlashSuccess(messages []string) {
	this.SetSession(FlashSuccess, messages)
}
func (this *BaseController) FlashInfo(messages []string) {
	this.SetSession(FlashInfo, messages)
}
