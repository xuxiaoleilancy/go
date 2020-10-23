package controllers

import (
	"webproject/rsystem/models"

	"github.com/astaxie/beego"
)

type ViewProductController struct {
	beego.Controller
}

func (this *ViewProductController) ListViewProducts() {

	maps := models.DefaultViewProductList.SelectAll()

	res := struct{ ViewProducts []*models.ViewProduct }{maps}
	this.Data["json"] = res
	this.ServeJSON()

	defer this.Render()
}
