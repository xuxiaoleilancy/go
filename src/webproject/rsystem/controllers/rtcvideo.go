package controllers

import (
	"github.com/astaxie/beego"
)

type RTCVideo struct {
	beego.Controller
}

func (this *RTCVideo) Get() {
	this.TplName = "rtcvideo.html"
	this.Render()
}
