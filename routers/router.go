package routers

import (
	"dudu/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

// init
func init() {
	// ---------------> rest posts -------------- //
	beego.CtrlPost("v1/user/register", (*controllers.UserController).Register)
	beego.CtrlGet("v1/user/login", (*controllers.UserController).Login)
	beego.CtrlGet("v1/user/getUserInfo", (*controllers.UserController).GetUserInfo)
	beego.CtrlGet("v1/user/logout", (*controllers.UserController).Logout)
}
