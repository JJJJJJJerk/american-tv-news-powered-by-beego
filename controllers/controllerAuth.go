package controllers

import (
	"encoding/json"
	"fmt"
	"my_go_web/models"
	"net/http"
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	beego.Controller //集成beego controller

}
type WeibAuth2Response struct {
	Access_token string `json:"access_token"`
	Uid          uint   `json:"uid,string"`
}

//http://open.weibo.com/wiki/2/users/show
type WeiboUser struct {
	Id           uint   `json:"id"`
	Screen_name  string `json:"screen_name"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Avatar_large string `json:"avatar_large"`
}

//sign up
func (c *AuthController) GetRegister() {
	user := models.User{}
	//weibo auth2 回调
	weiboCode := c.GetString("code")
	if weiboCode != "" {
		resp, _ := http.PostForm("https://api.weibo.com/oauth2/access_token", url.Values{
			"client_id":     {WeiboAppId},
			"client_secret": {WeiboAppSecret},
			"grant_type":    {"authorization_code"},
			"code":          {weiboCode},
			"redirect_uri":  {"https://www.mojotv.cn/auth/register"},
		})
		defer resp.Body.Close()

		//解析json 获取token和uid

		var weiboResponseJson WeibAuth2Response
		json.NewDecoder(resp.Body).Decode(&weiboResponseJson)

		//logs.Debug("token struct:", weiboResponseJson)
		//logs.Debug("token uid and access token", weiboResponseJson.Uid, weiboResponseJson.Access_token)

		logs.Debug("first weibo id", weiboResponseJson.Uid)

		models.Gorm.Where("weibo_id = ?", weiboResponseJson.Uid).First(&user)
		// logs.Debug("weibo struct user:%v", user)
		// logs.Debug("weibo struct accesstoken:%v", weiboResponseJson)
		if user.ID > 0 {
			//用户已注册
			//直接登陆
			c.SetSession(AuthSessionName, user)
			c.Redirect("/", 303)
			return
		}
		//用户未注册     获取用户信息
		getURL := fmt.Sprintf("https://api.weibo.com/2/users/show.json?access_token=%s&uid=%d", weiboResponseJson.Access_token, weiboResponseJson.Uid)
		respInfo, _ := http.Get(getURL)
		defer respInfo.Body.Close()

		//解析json 获取token和uid
		var weiboUser WeiboUser
		json.NewDecoder(respInfo.Body).Decode(&weiboUser)
		//logs.Debug("weibo name info:%v", weiboUser.Name)
		//logs.Debug("weibo struct info:%v", weiboUser)
		logs.Debug("first weibo info id", weiboUser.Id)

		user.WeiboId = weiboUser.Id
		user.Name = weiboUser.Name
		user.AvatarImage = weiboUser.Avatar_large
		user.Email = fmt.Sprint("weibo_", weiboUser.Id, "@mojotv.cn")
	}
	c.Data["User"] = user
	c.Data["Xsrf"] = c.XSRFToken() //防止跨域
	c.TplName = "auth/register.html"
}

func (c *AuthController) PostRegister() {
	password := c.GetString("password")
	passwordConfirmed := c.GetString("password_confirmed")
	if password == "" || passwordConfirmed == "" || (password != passwordConfirmed) {
		c.Data["json"] = map[string]interface{}{"status": "error", "message": "两次输入的密码不相同,或者密码为空", "data": nil}
		c.ServeJSON()
		return

	}

	email := c.GetString("email")
	var isExistUser models.User
	models.Gorm.Where("email = ?", email).First(&isExistUser)
	if isExistUser.ID > 0 && isExistUser.WeiboId < 1 {
		beego.Warning("用户已近存在")
		c.Data["json"] = map[string]interface{}{"status": "error", "message": "email已经注册", "data": nil}
		c.ServeJSON()
		return
	}

	//hash password
	password_byte := []byte(password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password_byte, bcrypt.DefaultCost)
	//new struct from package
	//TODO:需要搞清楚 go语言的 pointer * &的用法
	isExistUser.Email = email
	isExistUser.Password = string(hashedPassword)
	isExistUser.Name = c.GetString("name")
	isExistUser.WeiboAvatar = c.GetString("avatar_image")
	wbId, _ := c.GetInt("weibo_id")
	logs.Debug("jerk:", wbId)
	isExistUser.WeiboId = uint(wbId)
	logs.Debug("jerk:", isExistUser.WeiboId)

	models.Gorm.Create(&isExistUser)
	if isExistUser.ID < 1 {
		beego.Critical("用户注册数据库添加失败")
		c.Data["json"] = map[string]interface{}{"status": "error", "message": "添加新用户失败", "data": nil}
		c.ServeJSON()
		return

	} else {
		c.SetSession(AuthSessionName, isExistUser)
		c.Data["json"] = map[string]interface{}{"status": "success", "message": "添加新用户成功", "data": nil}
		c.ServeJSON()
		return
	}
}

func (c *AuthController) PostLogin() {
	email := c.GetString("email")
	user := models.User{}
	models.Gorm.Table("users").Where("email = ?", email).First(&user)

	if user.ID < 1 {
		c.Data["json"] = map[string]interface{}{"status": "error", "message": "用户不存在", "data": nil}
		c.ServeJSON()
		return
	} else {
		//比较密码
		//string to []byte
		password := []byte(c.GetString("password"))
		//http://stackoverflow.com/questions/23259586/bcrypt-password-hashing-in-golang-compatible-with-node-js

		// Hashing the password with the default cost of 10  DefaultCost int = 10
		//laravel bcrypt /Library/WebServer/Documents/estate/vendor/laravel/framework/src/Illuminate/Hashing/BcryptHasher.php
		// Comparing the password with the hash
		db_hashed_password := []byte(user.Password)
		err := bcrypt.CompareHashAndPassword(db_hashed_password, password)
		if err == nil { // nil means it is a match
			//设置登陆session info
			c.SetSession(AuthSessionName, user)
			c.Data["json"] = map[string]interface{}{"status": "success", "message": "用户登陆成功", "data": nil}
			c.ServeJSON()
			return
		} else {
			c.Data["json"] = map[string]interface{}{"status": "error", "message": "密码错误", "data": nil}
			c.ServeJSON()
			return
		}
	}
}

// //find password 填写email
func (c *AuthController) GetResetPassword() {

}
func (c *AuthController) PostResetPassword() {
	//获取email地址 发送邮件

}

// //注销
func (c *AuthController) GetLogout() {
	c.DelSession(AuthSessionName)
	c.Redirect("/", 302)
}

func (c *AuthController) ToutiaoAd() {
	c.Redirect("/", 301)

	// c.Data["Title"] = "mojoTV美剧|最新最快最热的美剧周边资讯"
	// c.Data["Keyword"] = "mojoTV美剧,轻松学英语,欧美美剧资讯,国外搞笑小视频,"
	// c.Data["Description"] = "mojoTV美剧|提供最新最热最快最热的美剧资讯,海量英语学习资源,欧美搞笑有创意的短视频gif动图,海量美剧双语原生字幕,这里是英语爱好者的乐园,让每一个人都爱上学习英语"
	// c.Layout = "layout/base_index.html"
	// c.TplName = "home/toutiaoAd.html"
}
