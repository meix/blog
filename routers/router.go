package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.HomeController{}, "get:Index")

	//  404 页面
	beego.Router("/404.html", &controllers.ApplicationController{}, "*:Go404")

	// 文章详情
	beego.Router("/article/:id", &controllers.ArticleController{}, "get:Show")

	// 登陆页面 (get;post)
	beego.Router("/login", &controllers.LoginUserController{})

	// 注册页面
	beego.Router("/signup", &controllers.SignupUserController{})

	// 关于
	beego.Router("/user", &controllers.UserController{}, "get:Index")

	// beego.Router("/", &controllers.ListArticleController{})

	// beego.Router("/article", &controllers.ListArticleController{})
	// beego.Router("/article/:id", &controllers.ShowArticleController{})

	// beego.Router("/login", &controllers.LoginUserController{})
	// beego.Router("/logout", &controllers.LogoutUserController{})

	// beego.Router("/article/add", &controllers.AddArticleController{})
	// beego.Router("/article/edit/:id", &controllers.EditArticleController{})

	// beego.Router("/comment/add", &controllers.AddCommentController{})
	// beego.Router("/comment/edit/status", &controllers.EditCommentController{})

	// beego.Router("/album", &controllers.ListAlbumController{})
	// beego.Router("/album/upload", &controllers.UploadAlbumController{})
	// beego.Router("/album/edit", &controllers.EditAlbumController{})

	// beego.Router("/about", &controllers.AboutUserController{})

	// beego.Router("/uploadmulti", &controllers.UploadMultiController{})
	// beego.Router("/upload", &controllers.UploadController{})
}
