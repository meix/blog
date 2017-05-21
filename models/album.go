package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	// "time"
)

type Album struct {
	Id       int
	Title    string
	Picture  string
	Keywords string
	Summary  string
	Created  int64
	Viewnum  int64
	Status   int
}

func (this *Album) TableName() string {
	return "album"
}

func init() {
	orm.RegisterModel(new(Album))
}
