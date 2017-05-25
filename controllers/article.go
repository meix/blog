package controllers

import (
	. "blog/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"strconv"
)

type ArticleController struct {
	ApplicationController
}

// 博客详情
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

// 创建博客
func (this *ArticleController) Get() {
	this.RequireUser()

	this.Layout = "layout/application.tpl"
	this.TplName = "article/new.tpl"
}

func (this *ArticleController) Post() {
	this.RequireUser()
	// user_id := this.GetSession("user_id").(int)

	title := this.GetString("title")
	keywords := this.GetString("keywords")
	summary := this.GetString("summary")
	content := this.GetString("content")

	ArticleParams := make(map[string]string)
	ArticleParams["title"] = title
	ArticleParams["keywords"] = keywords
	ArticleParams["summary"] = summary
	ArticleParams["content"] = content

	CreateArticle(ArticleParams)

	// if CreateArticle(ArticleParams, user_id) {
	this.Redirect("/article/personal", 301)
	// } else {
	// 	this.Redirect("/article", 301)
	// }
}

// 私有的博客
func (this *ArticleController) Personal() {
	this.RequireUser()

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
	this.TplName = "article/personal.tpl"

}
