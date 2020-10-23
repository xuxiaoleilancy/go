package routers

import (
	"webproject/rsystem/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/todo/", &controllers.TodoController{})
	beego.Router("/todo/task/", &controllers.TaskController{}, "get:ListTasks;post:NewTask")
	beego.Router("/todo/task/:id:int", &controllers.TaskController{}, "get:GetTask;put:UpdateTask")

	beego.Router("/rautomatic", &controllers.RAutomaticCenter{})
	beego.Router("/rautomatic/product/", &controllers.ProductController{}, "get:ListProducts;post:NewProduct")
	beego.Router("/rautomatic/product/:id:int", &controllers.ProductController{}, "get:GetProduct;put:UpdateProduct")

	beego.Router("/rautomatic/viewproduct", &controllers.ViewProductController{}, "get:ListViewProducts")

	beego.Router("/rtcvideo", &controllers.RTCVideo{})
	beego.Router("/rdoc", &controllers.RDoc{})

	beego.Router("/rtest", &controllers.RTestController{})
	beego.Router("/rtest/cmds", &controllers.CmdController{})

}
