package models

import (
	// "fmt"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	// "reflect"
	"time"
)

type User struct {
	Id         int
	Email      string    // 邮箱
	Password   string    // 密码
	Avatar     string    // 头像
	Sex        string    // 1:男， 2：女
	Phone      int64     // 手机号
	Address    string    // 住址
	Education  string    // 教育
	Name       string    // 真实姓名
	Introduce  string    // 个人简介
	LoginIp    string    // 登录ip
	CreateTime time.Time // 创建时间

	Articles []*Article `orm:"reverse(many)"` // fk 的反向关系
	Comments []*Comment `orm:"reverse(many)"` // fk 的反向关系
}

func (this *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}

// 创建用户
func CreateUser(params map[string]string) (int64, error) {
	user_password, err := GetPassword(params["password"])
	if err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	o.Using("default")
	user := new(User)
	user.Email = params["email"]
	user.Password = user_password
	user.Name = params["name"]
	user.CreateTime = time.Now()
	id, err := o.Insert(user)
	return id, err
}

// 用户登陆验证
func LoginUser(email, password string) (User, error) {
	var user User
	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("email", email).One(&user)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		return user, err
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		return user, err
	}
	// 比较密码
	err = bcrypt.CompareHashAndPassword([]uint8(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, err
}

// 获取当前用户
func GetUser(id int) (User, error) {
	var user User
	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("id", id).One(&user)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		return user, err
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		return user, err
	}
	return user, err
}

// 获取密码
func GetPassword(psw string) (string, error) {
	password := []byte(psw)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return string(hashedPassword), err
}
