package controllers

import (
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller //集成beego controller
	IsLogin          bool
	UserUserId       int64
	UserUsername     string
	UserAvatar       string
}

func (this *BaseController) Prepare() {
	//判断用户数是否已近登陆
	userLogin := this.GetSession("userLogin")
	if userLogin == nil {
		this.IsLogin = false
	} else {
		this.IsLogin = true
		tmp := strings.Split((this.GetSession("userLogin")).(string), "||")

		userid, _ := strconv.Atoi(tmp[0])
		longid := int64(userid)
		this.Data["LoginUserid"] = longid
		this.Data["LoginUsername"] = tmp[1]
		this.Data["LoginAvatar"] = tmp[2]

		this.UserUserId = longid
		this.UserUsername = tmp[1]
		this.UserAvatar = tmp[2]

		this.Data["PermissionModel"] = this.GetSession("userPermissionModel")
		this.Data["PermissionModelc"] = this.GetSession("userPermissionModelc")
		//TODO:显示消息

	}
	this.Data["IsLogin"] = this.IsLogin
}
