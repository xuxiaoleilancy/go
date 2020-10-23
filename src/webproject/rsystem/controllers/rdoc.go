package controllers

import (
	"github.com/astaxie/beego"
)

type RDoc struct {
	beego.Controller
}

func (this *RDoc) Get() {
	this.TplName = "dev_doc_paassdk_windows.html"
	this.Render()
}
