package models

import (
	"dudu/util"
	"fmt"
	"time"

	// "strconv"
	// "time"
	orm "github.com/beego/beego/v2/client/orm"
)

// 一个未开通应用新注册的用户的信息包括
// name
// telephone
// password
// email
// profile
// createdStamp
// updatedStamp

type User struct {
	Id           uint64    `orm:"column(id)"`
	Name         string    `orm:"column(name)"`
	Telephone    string    `orm:"column(telephone)"`
	PassWord     string    `orm:"column(password)"`
	Email        string    `orm:"column(email)"`
	CreatedStamp time.Time `orm:"column(createdStamp);auto_now_add;type(date);null"`
	UpdatedStamp time.Time `orm:"column(updatedStamp);auto_now;type(date);null"`
	Status       int       `orm:"default(1);description(这是状态字段)"`
}

func (u *User) TableName() string {
	return "user"
}

// Insert
//
//	@receiver u
//	@param data
//	@return int16
//	@return error
//
// 插入用户数据
func (u *User) Insert(data User) (int16, error) {
	user := new(User)
	user.Name = data.Name
	user.Telephone = data.Telephone
	user.Email = data.Email

	pwd, _ := util.GetPwd(string(data.PassWord))
	user.PassWord = string(pwd)

	ret, error := orm.NewOrm().Insert(user)
	return int16(ret), error
}

func (u *User) GetUser(email string) User {
	usr := User{}

	return usr
}

func SliceToFlat(data []string) string {
	length := len(data)

	if length == 0 {
		return ""
	}

	result := ""
	for k, v := range data {
		result += v
		if k < (length - 1) {
			result += ","
		}
	}

	return result
}

func (u *User) ValidUserInfo(data User) bool {
	// valid := validation.Validation{}
	return true
}

func (u *User) HasUser(data User) User {

	// 创建orm对象
	o := orm.NewOrm()

	// 获取 QuerySeter 对象，并设置表名orders
	qs := o.QueryTable("user")

	// 定义保存查询结果的变量
	user := User{}

	// 使用QuerySeter 对象构造查询条件，并执行查询。
	err := qs.Filter("telephone", data.Telephone).One(&user)

	if err == orm.ErrMultiRows {
		// 多条的时候报错
		fmt.Printf("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		fmt.Printf("Not row found")
	}

	return user
}

func (u *User) FindUserByEmail(data User) User {

	// 创建orm对象
	o := orm.NewOrm()

	// 获取 QuerySeter 对象，并设置表名orders
	qs := o.QueryTable("user")

	// 定义保存查询结果的变量
	user := User{}

	// 使用QuerySeter 对象构造查询条件，并执行查询。
	err := qs.Filter("email", data.Email).One(&user)

	if err == orm.ErrMultiRows {
		// 多条的时候报错
		fmt.Printf("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		fmt.Printf("Not row found")
	}

	return user
}

func init() {
	fmt.Println("user model init")
}
