package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Article struct {
	Id         int
	Title      string    // 标题
	Keywords   string    // 关键字
	Summary    string    // 简介
	Content    string    // 内容
	CreateTime time.Time // 创建时间
	Viewnum    int       // 浏览量
	Status     int       // 状态

	User     *User      `orm:"rel(fk)"`       // RelForeignKey relation
	Comments []*Comment `orm:"reverse(many)"` // fk 的反向关系
}

func (this *Article) TableName() string {
	return "article"
}

func init() {
	orm.RegisterModel(new(Article))
}

func (this *Article) GetArticle(id int) *Article {
	this.Id = id
	return this
}

func CreateArticle(params map[string]string) bool {
	o := orm.NewOrm()
	o.Using("default")
	article := new(Article)

	article.Title = params["title"]
	article.Keywords = params["keywords"]
	article.Summary = params["summary"]
	article.Content = params["content"]
	article.CreateTime = time.Now()
	article.User = 1

	_, err := o.Insert(article)

	fmt.Printf("错误是", err)

	if err != nil {
		return false
	}
	return true
}
