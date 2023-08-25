package models

import (
	"fmt"
	"time"

	// "strconv"
	// "time"
	orm "github.com/beego/beego/v2/client/orm"
	"golang.org/x/crypto/bcrypt"
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
	Id           uint64    `orm:"column(id);orm:auto"`
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

// GetPwd 给密码加密
func GetPwd(pwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return hash, err
}

// ComparePwd 比对密码
func ComparePwd(pwd1 string, pwd2 string) bool {
	// Returns true on success, pwd1 is for the database.
	err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
	if err != nil {
		return false
	} else {
		return true
	}
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

	pwd, _ := GetPwd(data.PassWord)
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

func (u *User) VerifyUser(data User) bool {
	fmt.Println("---------> 查询用户的信息", data.Telephone)
	qb, _ := orm.NewQueryBuilder("mysql")

	var local []string

	qb.Select("telephone").
		From("user").
		Where("telephone = ?").
		Limit(10).Offset(0)

	sql := qb.String()

	o := orm.NewOrm()
	tel := data.Telephone[:]
	err, _ := o.Raw(sql, tel).QueryRows(&local)

	fmt.Println("---------> 查询到的数据行数", err, local)
	if cap(local) > 0 {
		return false
	} else {
		return true
	}
}

func init() {
	fmt.Println("user model init")
}
