package controllers

import (
	"my_go_web/models"

	"github.com/astaxie/beego"

	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	BaseController
}

//sign up
func (c *AuthController) GetRegister() {
	c.TplName = "auth/register.html"
}

func (c *AuthController) PostRegister() {
	password := c.GetString("password")
	passwordConfirmed := c.GetString("password_confirmed")
	if password == "" || passwordConfirmed == "" || (password != passwordConfirmed) {
		c.JsonRetrun("error", "两次输入的密码不相同,或者密码为空", nil)
		return

	}

	email := c.GetString("email")
	isExistUser := models.User{}
	models.Gorm.Where("email = ?", email).First(&isExistUser)
	if isExistUser.ID > 0 {
		beego.Warning("用户已近存在")
		c.JsonRetrun("error", "email已经注册", nil)
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
		c.JsonRetrun("error", "添加新用户失败", nil)
		return

	} else {
		c.SetSession("loginInfo", isExistUser)
		c.JsonRetrun("success", "添加新用户成功", nil)
		return

	}
}

func (c *AuthController) PostLogin() {
	email := c.GetString("email")
	user := models.User{}
	models.Gorm.Where("email = ?", email).First(&user)

	if user.ID < 1 {
		c.JsonRetrun("error", "用户不存在", nil)
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
			c.JsonRetrun("success", "用户登陆成功", user)
		} else {
			c.JsonRetrun("error", "密码错误", nil)
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
	session := c.GetSession("loginInfo")
	if session != nil {
		c.DelSession("loginInfo")
	}
	c.Redirect("/", 302)
}
