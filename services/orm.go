package services

import "fmt"
import orm "github.com/beego/beego/v2/client/orm"
import _ "github.com/go-sql-driver/mysql"
import "github.com/beego/beego/v2/core/logs"
import "dudu/models"

func ormInit() {
	logs.Info("开始注册数据库模型")
	// need to register models in init
	orm.RegisterModel(new(models.User))
	// need to register default database
	orm.RegisterDataBase("default", "mysql", "root:12345678@tcp(127.0.0.1:3306)/beegoSSO?charset=utf8")

	orm.SetMaxOpenConns("default", 30);
}


func startOrm () {
	// automatically build table
	orm.RunSyncdb("default", false, true)

	// create orm object
	o := orm.NewOrm()

	// data
	user := new(models.User)
	user.Name = "mike"

	// insert data
	o.Insert(user)
}


func init() {
	fmt.Println("orm init");

	ormInit()
	startOrm()
}