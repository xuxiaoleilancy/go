package controllers

import (
	"github.com/astaxie/beego"
)

type Todo struct {
	beego.Controller
}

func (this *Todo) Get() {
	this.TplName = "todo.html"
	this.Render()
}
