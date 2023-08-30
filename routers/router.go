package routers

import (
	"dudu/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

// init
func init() {
	// 用户
	// beego.CtrlGet("v1/user", (*controllers.UserController).Register)

	beego.CtrlGet("/", (*controllers.HomeController).GetTemplate)

	// 登录静态页面
	beego.CtrlGet("/login", (*controllers.HomeController).GetTemplate)

	// 注册静态页面
	// beego.CtrlGet("register", (*controllers.HomeController).GetTemplate)

	// user register
	beego.CtrlPost("v1/user/register", (*controllers.UserController).Register)

	beego.CtrlGet("v1/user/login", (*controllers.UserController).Login)

	//用户信息
	beego.CtrlGet("v1/user/getUserInfo", (*controllers.UserController).GetUserInfo)

	//用户登出
	beego.CtrlGet("v1/user/logout", (*controllers.UserController).Logout)
}
