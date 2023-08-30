package main

import (
	_ "dudu/models"
	_ "dudu/routers"
	_ "dudu/services"

	"github.com/beego/beego/v2/client/orm"
	server "github.com/beego/beego/v2/server/web"
)

func main() {
	if server.BConfig.RunMode == "dev" {
		orm.Debug = true
		server.BConfig.Listen.AdminAddr = "192.168.0.188"
		server.BConfig.Listen.AdminPort = 9999
	}

	server.Run()
}
