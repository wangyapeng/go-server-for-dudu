package controllers

import (

	// "encoding/json"

	"encoding/json"
	"fmt"

	"dudu/models"

	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)

	fmt.Println(user)
	u.SetData("succees")
	u.Data["data"] = "delete success!"
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) Register() {

	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)

	fmt.Println(user)
	u.SetData("succees")
	u.Data["data"] = "delete success!"
	u.ServeJSON()
}

func init() {
	models.Test()
	models.Test3()
}
