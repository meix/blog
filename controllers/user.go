package controllers

import (
	. "blog/models"
	"fmt"
)

// 用户 登陆
type LoginUserController struct {
	// beego.Controller
	ApplicationController
}

func (this *LoginUserController) Get() {
	this.Layout = "layout/application.tpl"
	this.TplName = "user/login.tpl"
}

func (this *LoginUserController) Post() {
	email := this.GetString("email")
	password := this.GetString("password")
	user, err := LoginUser(email, password)
	if err != nil {
		fmt.Println("邮箱或密码错误,请重新输入：", err)
		this.Redirect("/login", 302)
	} else {
		this.SetSession("user_id", user.Id)
		this.Redirect("/", 302)
	}
}

// 用户 注册
type SignupUserController struct {
	ApplicationController
}

func (this *SignupUserController) Get() {
	this.Layout = "layout/application.tpl"
	this.TplName = "user/signup.tpl"
}

func (this *SignupUserController) Post() {
	email := this.GetString("email")
	name := this.GetString("name")
	password := this.GetString("password")
	password_confirmation := this.GetString("password_confirmation")

	if password != password_confirmation {
		fmt.Println("密码与重复密码不相等！")
		this.Redirect("/signup", 302)
	}

	UserParams := make(map[string]string)
	UserParams["email"] = email
	UserParams["name"] = name
	UserParams["password"] = password

	user_id, err := CreateUser(UserParams)
	if err != nil {
		fmt.Println("注册错误：", err)
		this.Redirect("/signup", 302)
	} else {
		this.SetSession("user_id", int(user_id))
		this.Redirect("/", 302)
	}
}

type UserController struct {
	ApplicationController
}

// 退出 登陆
func (this *UserController) Signout() {
	user_id := this.GetSession("user_id")
	if user_id != nil {
		this.DelSession("user_id")
	}
	this.Redirect("/", 302)
}

// 个人页面
func (this *UserController) Index() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"

	this.Layout = "layout/application.tpl"
	this.TplName = "user/index.tpl"
}
