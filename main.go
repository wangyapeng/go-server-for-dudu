package main

import (
	_ "dudu/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		// beego.BConfig.WebConfig.DirectoryIndex = false
		// beego.BConfig.WebConfig.ViewsPath = "views"
		// beego.BConfig.WebConfig.StaticDir["/static"] = "static"
	}
	beego.Run()
}
