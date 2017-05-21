package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	// "time"
)

type Article struct {
	Id       int
	Title    string
	Uri      string
	Keywords string
	Summary  string
	Content  string
	Author   string
	Created  int64
	Viewnum  int
	Status   int

	Comment []*Comment `orm:"reverse(many)"` // fk 的反向关系
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
