package controllers

import (
	"html/template"
	"my_go_web/models"

	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	beego.Controller
}

//sign up
func (c *AuthController) GetRegister() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())

	c.TplName = "auth/register.html"
}
func (c *AuthController) PostRegister() {
	email := c.GetString("email")
	password := c.GetString("password")
	password_comfirmed := c.GetString("password_comfirmed")
	if password == password_comfirmed {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "两次输入的密码不相同"}
		c.ServeJSON()
	}
	//hash password
	password_byte := []byte(password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password_byte, bcrypt.DefaultCost)
	//new struct from package
	//TODO:需要搞清楚 go语言的 pointer * &的用法
	var user_model models.Users
	user_model.Email = email
	user_model.Password = string(hashedPassword)
	_, err := models.AddUsers(&user_model)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "创建用户错误"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "注册成功", "data": user_model}
	}
	c.ServeJSON()
}

//sign in
func (c *AuthController) GetLogin() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "auth/login.html"

}
func (c *AuthController) PostLogin() {
	email := c.GetString("email")
	user, err := models.GetUsersByEmail(email)

	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "用户名不存在"}
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
			c.SetSession("loginInfo", user)
			c.Data["json"] = map[string]interface{}{"code": 1, "message": "登陆成功"}
		}
	}
	c.ServeJSON()

}

//find password 填写email
func (c *AuthController) GetResetPassword() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())

}
func (c *AuthController) PostResetPassword() {
	//获取email地址 发送邮件

}

//注销
func (c *AuthController) GetLogout() {
	session := c.GetSession("loginInfo")
	if session == nil {
		c.DelSession("loginInfo")
	}
	c.Redirect("/", 302)
}
