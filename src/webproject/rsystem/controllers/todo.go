package controllers

import (
	"github.com/astaxie/beego"
)

type TodoController struct {
	beego.Controller
}

func (this *TodoController) Get() {
	this.TplName = "todo.html"
	this.Render()
}
