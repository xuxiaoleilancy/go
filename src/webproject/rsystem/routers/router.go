package routers

import (
	"webproject/rsystem/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/todo", &controllers.Todo{})
	beego.Router("/todo/task/", &controllers.TaskController{}, "get:ListTasks;post:NewTask")
	beego.Router("/todo/task/:id:int", &controllers.TaskController{}, "get:GetTask;put:UpdateTask")
}
