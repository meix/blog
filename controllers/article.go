package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/utils/pagination"
	// . "github.com/lock-upme/beegoblog/models"
	. "blog/models"
	"fmt"
	"strconv"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Show() {

	idstr := this.Ctx.Input.Param(":id")
	intid, _ := strconv.Atoi(idstr)

	o := orm.NewOrm()
	article := &Article{}
	err := o.QueryTable("article").Filter("Id", intid).RelatedSel().One(article)

	if err == orm.ErrMultiRows {
		// 多条的时候报错
		fmt.Printf("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		fmt.Printf("Not row found")
	}

	this.Data["article"] = article
	this.Layout = "layout/application.tpl"
	this.TplName = "article/show.tpl"
}
