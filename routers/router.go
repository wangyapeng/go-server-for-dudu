package routers

import (
	"dudu/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

// init
func init() {
	// 用户
	// beego.CtrlGet("v1/user", (*controllers.UserController).Register)

	// user register
	beego.CtrlPost("v1/user/register", (*controllers.UserController).Register)

	beego.CtrlGet("v1/user/login", (*controllers.UserController).Login)

	beego.CtrlGet("v1/user/logout", (*controllers.UserController).Logout)
}
