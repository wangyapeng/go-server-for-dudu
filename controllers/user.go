package controllers

import (

	// "encoding/json"

	"dudu/services"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

var userService = services.UserService{}

// user register
// return user.id
func (u *UserController) Register() {
	id := userService.CreateNewUser(u.Ctx.Input.RequestBody)
	if id == 0 {
		u.Ctx.Output.Status = 500
		u.Ctx.Resp(map[string]string{"errorMsg": "用户已存在", "errorCode": "5001200"})
	} else {
		u.Data["json"] = id
		u.Ctx.Resp(map[string]any{"data": id, "success": true})
	}
}

func (u *UserController) Login() {
	// res := userService.UserLogin(u.Ctx.Input.RequestBody)
	email := u.Ctx.Input.Query("email")
	password := u.Ctx.Input.Query("password")
	res := userService.UserLogin(email, password)
	fmt.Println(res)

	u.Ctx.Resp("lgoin")
}

func (u *UserController) Logout() {
	u.Ctx.Resp("logout")
}

func init() {

}
