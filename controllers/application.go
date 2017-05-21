package controllers

import (
	. "blog/models"
	"github.com/astaxie/beego"
)

type ApplicationController struct {
	beego.Controller
	isLogin bool
}

func (this *ApplicationController) Prepare() {
	user_id := this.GetSession("user_id")
	if user_id == nil {
		this.isLogin = false
	} else {
		this.isLogin = true
		user, err := GetUser(1)
		if err == nil {
			this.Data["user"] = user
		}
	}
	this.Data["isLogin"] = this.isLogin

}

func (this *ApplicationController) Go404() {
	this.Layout = "layout/application.tpl"
	this.TplName = "404.tpl"
	return
}
