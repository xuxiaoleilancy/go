package controllers

import (
	"fmt"

	"github.com/astaxie/beego"

	"webproject/rsystem/models"
)

type CmdController struct {
	beego.Controller
}

func (this *CmdController) Get() {
	maps := models.DefaultRTestList.SelectAll()

	fmt.Println(maps)
	res := struct{ Cmds []*models.RTest }{maps}
	this.Data["json"] = res
	this.ServeJSON()

	defer this.Render()
}
