package routers

import (
	"dudu/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

// init
func init() {
	// 用户
	beego.CtrlGet("v1/user", (*controllers.UserController).Register)

	// user register
	beego.CtrlPost("v1/user/register", (*controllers.UserController).Post)
}
