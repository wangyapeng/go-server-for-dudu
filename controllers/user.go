package controllers

import (

	// "encoding/json"

	"dudu/models"
	"dudu/services"
	"dudu/util"
	"fmt"
	"strconv"

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
		u.ServeJSON()
	} else {
		u.Data["json"] = id
		u.Ctx.Resp(map[string]any{"data": id, "success": true})
		u.ServeJSON()
	}
}

func (u *UserController) Login() {
	email := u.Ctx.Input.Query("email")
	telephone := u.Ctx.Input.Query("telephone")
	password := u.Ctx.Input.Query("password")

	fmt.Println(telephone, password)
	id, err := userService.UserLogin(email, telephone, password)

	if err != nil {
		u.Ctx.Output.Status = 500
		u.Ctx.Resp(map[string]string{"errorMsg": err.Error(), "errorCode": "5001200"})
	} else {
		token, _ := util.GenerateToken()

		u.Ctx.Output.EnableGzip = true
		u.Ctx.SetCookie("login", "true", 1000*60*60*24, "/", ".dq.com")
		u.Ctx.SetCookie("userToken", token, 1000*60*60*24, "/", ".dq.com")
		u.Ctx.SetCookie("userId", strconv.Itoa(int(id)), 1000*60*60*24, "/", ".dq.com")
		u.Ctx.Resp(map[string]string{"token": token})
		u.ServeJSON()
	}
}

func (u *UserController) GetUserInfo() {
	id := u.Ctx.Input.Query("id")
	user := models.User{}
	intId, _ := strconv.ParseInt(id, 10, 64)
	res := user.FindUserById(intId)
	if id == "" {
		u.Ctx.Output.Status = 500
		u.Ctx.Resp(map[string]string{"errorMsg": "参数错误", "errorCode": "5001200"})
		u.ServeJSON()
	}

	if res.Id == 0 {
		u.Ctx.Output.Status = 500
		u.Ctx.Resp(map[string]string{"errorMsg": "未查询到信息", "errorCode": "5001201"})
		u.ServeJSON()
	}

	u.Ctx.Resp(res)
	u.ServeJSON()
}

func (u *UserController) Logout() {
	u.Ctx.Resp("logout")
}

func init() {

}
