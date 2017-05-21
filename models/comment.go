package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	// "time"
)

type Comment struct {
	Id       int
	nickname string
	Content  string
	Created  int64
	Status   int

	Article *Article `orm:"rel(fk)"` // RelForeignKey relation
}

func (this *Comment) TableName() string {
	return "comment"
}

func init() {
	orm.RegisterModel(new(Comment))
}
