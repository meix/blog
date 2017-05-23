package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.HomeController{}, "get:Index")

	//  404 页面
	beego.Router("/404.html", &controllers.ApplicationController{}, "*:Go404")

	// 登陆页面 (get;post)
	beego.Router("/login", &controllers.LoginUserController{})

	// 注册页面
	beego.Router("/signup", &controllers.SignupUserController{})

	// 退出
	beego.Router("/signout", &controllers.UserController{}, "get:Signout")

	// 关于
	beego.Router("/user", &controllers.UserController{}, "get:Index")

	// 文章详情
	beego.Router("/article/:id", &controllers.ArticleController{}, "get:Show")

	// 个人博客列表
	beego.Router("/article", &controllers.ArticleController{})

	// 个人博客列表
	beego.Router("/article/personal", &controllers.ArticleController{}, "get:Personal")

}
