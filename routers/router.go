package routers

import (
	"dudu/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func intercepter() {

	// beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	// 	AllowAllOrigins:  true,
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
	// 	AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	// 	ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	// 	AllowCredentials: true,
	// }))
}

// init
func init() {

	intercepter()

	// 用户
	// beego.CtrlGet("v1/user", (*controllers.UserController).Register)

	beego.CtrlGet("/", (*controllers.HomeController).GetTemplate)

	// 登录静态页面
	beego.CtrlGet("/login", (*controllers.HomeController).GetTemplate)

	// ---------------> rest posts -------------- //
	beego.CtrlPost("v1/user/register", (*controllers.UserController).Register)
	beego.CtrlGet("v1/user/login", (*controllers.UserController).Login)
	beego.CtrlGet("v1/user/getUserInfo", (*controllers.UserController).GetUserInfo)
	beego.CtrlGet("v1/user/logout", (*controllers.UserController).Logout)
}
