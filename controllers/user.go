package controllers

import (

	// "encoding/json"

	"dudu/services"
	"dudu/util"
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
	id, err := userService.CreateNewUser(u.Ctx.Input.RequestBody)

	if err != nil {
		u.Ctx.Output.Status = 500
		u.Ctx.Resp(map[string]string{"errorMsg": err.Error(), "errorCode": "5001200"})
	} else {
		u.Data["json"] = id
		u.Ctx.Resp(map[string]any{"data": id, "success": true})
	}
}

func (u *UserController) Login() {
	email := u.Ctx.Input.Query("email")
	telephone := u.Ctx.Input.Query("telephone")
	password := u.Ctx.Input.Query("password")
	_, err := userService.UserLogin(email, telephone, password)

	if err != nil {
		u.Ctx.Output.Status = 500
		u.Ctx.Resp(map[string]string{"errorMsg": err.Error(), "errorCode": "5001200"})
	} else {
		token, _ := util.GenerateToken()
		fmt.Println("------- controller", token)
		u.Ctx.Resp(token)
	}
}

func (u *UserController) Logout() {
	u.Ctx.Resp("logout")
}

func init() {

}
