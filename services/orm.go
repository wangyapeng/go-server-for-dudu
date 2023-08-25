package services

import (
	"dudu/models"

	orm "github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
)

func ormInit() {
	logs.Info("开始注册数据库模型")
	// need to register models in init
	orm.RegisterModel(new(models.User))
	// need to register default database
	orm.RegisterDataBase("default", "mysql", "root:12345678@tcp(127.0.0.1:3306)/beegoSSO?charset=utf8")

	orm.SetMaxOpenConns("default", 30)
}

func init() {
	ormInit()

	orm.RunSyncdb("default", false, true)

}
