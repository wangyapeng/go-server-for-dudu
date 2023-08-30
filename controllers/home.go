package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type HomeController struct {
	beego.Controller
}

func (u *HomeController) GetTemplate() {

	u.Ctx.Output.Body([]byte("hello, go"))
}

func init() {

}
