package controllers

import (
	"github.com/astaxie/beego"
)

type RTestController struct {
	beego.Controller
}

func (this *RTestController) Get() {
	this.TplName = "rtest.html"
	this.Render()
}
