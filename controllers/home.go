package controllers

import (
	"fmt"
	"os"

	beego "github.com/beego/beego/v2/server/web"
)

type HomeController struct {
	beego.Controller
}

func (u *HomeController) GetTemplate() {
	file, err := os.ReadFile("./www/dist/index.html")
	if err != nil {
		fmt.Println(err)
	}
	u.Ctx.Output.Body(file)
}

func init() {

}
