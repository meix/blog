package controllers

import (
	_ "blog/helpers"
	. "blog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
)

type HomeController struct {
	// beego.Controller
	ApplicationController
}

func (this *HomeController) Index() {

	page, page_err := this.GetInt("p")
	if page_err != nil {
		page = 1
	}

	var articles []*Article
	o := orm.NewOrm()
	qs := o.QueryTable("article").Filter("status", 1)

	cnt, _ := qs.Count()

	offset, _ := beego.AppConfig.Int("pageoffset")
	paginator := pagination.SetPaginator(this.Ctx, offset, cnt)

	start := (page - 1) * offset
	qs.Limit(offset, start).All(&articles)

	this.Data["articles"] = articles
	this.Data["paginator"] = paginator

	this.Layout = "layout/application.tpl"
	this.TplName = "home/index.tpl"
}
