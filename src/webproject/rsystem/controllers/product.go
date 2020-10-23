package controllers

import (
	"encoding/json"

	"strconv"
	"webproject/rsystem/models"

	"github.com/astaxie/beego"
)

type ProductController struct {
	beego.Controller
}

func (this *ProductController) ListProducts() {

	maps := models.DefaultProductList.SelectAll()

	res := struct{ Products []*models.Product }{maps}
	this.Data["json"] = res
	this.ServeJSON()

	defer this.Render()
}

func (this *ProductController) NewProduct() {

	req := struct {
		Name string
	}{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	node, err := models.DefaultJenkinsNodeList.Find(1)
	if err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("JenkinsNode not found"))
		return
	}
	p := models.Product{Name: req.Name, JenkinsNode: node}
	models.DefaultProductList.Insert(&p)

	this.Render()
}

func (this *ProductController) GetProduct() {
	id := this.Ctx.Input.Param(":id")
	beego.Info("Product is ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	p, err := models.DefaultProductList.Find(intid)

	defer this.Render()
	if err != nil {

		beego.Info("Found", err)
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("Product not found"))

	} else {

		this.Data["json"] = p
		this.ServeJSON()
	}
}

func (this *ProductController) UpdateProduct() {
	id := this.Ctx.Input.Param(":id")
	beego.Info("Product is ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	var p models.Product
	defer this.Render()
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &p); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	if p.Id != intid {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("inconsistent Product IDs"))
		return
	}
	models.DefaultProductList.Update(p)
}
