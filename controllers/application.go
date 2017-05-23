package controllers

import (
	. "blog/models"

	"github.com/astaxie/beego"
)

type ApplicationController struct {
	beego.Controller
	isLogin      bool
	current_user User
}

func (this *ApplicationController) Prepare() {
	user_id := this.GetSession("user_id")
	if user_id == nil {
		this.isLogin = false
	} else {
		this.isLogin = true
		user, err := GetUser(user_id.(int))
		if err == nil {
			this.current_user = user
			this.Data["currentUser"] = user
		}
	}
	this.Data["isLogin"] = this.isLogin
}

// 如果没有登陆 跳转到登陆页面
func (this *ApplicationController) RequireUser() {
	if !this.isLogin {
		this.Redirect("/login", 302)
	}
}

func (this *ApplicationController) Go404() {
	this.Layout = "layout/application.tpl"
	this.TplName = "404.tpl"
	return
}
