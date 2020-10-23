package controllers

import (
	"github.com/astaxie/beego"
)

type RAutomaticCenter struct {
	beego.Controller
}

func (this *RAutomaticCenter) Get() {
	this.TplName = "viewproduct.html"
	this.Render()
}
