package models

import (
	// "fmt"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	// "reflect"
)

type User struct {
	Id          int
	Email       string
	UserProfile *UserProfile `orm:"rel(one)"`
	Password    string
	Status      int
	Created     int64
	Changed     int64
}

type UserProfile struct {
	Id       int
	Realname string
	Sex      int
	Birth    string
	Email    string
	Phone    string
	Address  string
	Hobby    string
	Intro    string
	User     *User `orm:"reverse(one)"`
}

func (this *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User), new(UserProfile))
}

// 创建用户
func CreateUser(params map[string]string) interface{} {
	user_password, err := GetPassword(params["password"])
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	o.Using("default")

	user_profile := new(UserProfile)
	user_profile.Realname = params["name"]
	user_profile.Email = params["email"]

	user := new(User)
	user.UserProfile = user_profile
	user.Email = params["email"]
	user.Password = user_password

	o.Insert(user_profile)
	o.Insert(user)
	// fmt.Println(id)
	return err
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
