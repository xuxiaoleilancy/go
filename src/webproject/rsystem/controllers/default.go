package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	localIp := beego.AppConfig.DefaultString("localip", "127.0.0.1")
	c.Data["Website"] = localIp
	c.Data["Email"] = "xl.xu@suirui.com"
	c.TplName = "index.tpl"
}
