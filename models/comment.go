package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Comment struct {
	Id         int
	Content    string    // 内容
	Status     string    // 状态
	CreateTime time.Time // 创建时间

	Article *Article `orm:"rel(fk)"` // RelForeignKey relation
	User    *User    `orm:"rel(fk)"` // RelForeignKey relation
}

func (this *Comment) TableName() string {
	return "comment"
}

func init() {
	orm.RegisterModel(new(Comment))
}
