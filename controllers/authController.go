package controllers

import (
	"my_go_web/models"

	"github.com/astaxie/beego"

	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	beego.Controller //集成beego controller

}

//sign up
func (c *AuthController) GetRegister() {
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
	isExistUser := models.User{}
	models.Gorm.Where("email = ?", email).First(&isExistUser)
	if isExistUser.ID > 0 {
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
	models.Gorm.Create(&isExistUser)
	if isExistUser.ID < 1 {
		beego.Critical("用户注册数据库添加失败")
		c.Data["json"] = map[string]interface{}{"status": "error", "message": "添加新用户失败", "data": nil}
		c.ServeJSON()
		return

	} else {
		c.SetSession("loginInfo", isExistUser)
		c.Data["json"] = map[string]interface{}{"status": "success", "message": "添加新用户成功", "data": nil}
		c.ServeJSON()
		return

	}
}

func (c *AuthController) PostLogin() {
	email := c.GetString("email")
	user := models.User{}
	models.Gorm.Where("email = ?", email).First(&user)

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
