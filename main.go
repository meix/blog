package main

import (
	_ "blog/initial"
	_ "blog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"html/template"
	"net/http"
)

func main() {
	beego.ErrorHandler("404", page_not_found)
	beego.SetStaticPath("/static", "static")

	// 开启 ORM 调试模式
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, false)

	beego.Run()
}

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.tpl").ParseFiles("views/404.tpl")

	data := make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}
