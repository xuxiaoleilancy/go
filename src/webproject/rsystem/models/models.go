package models

import (
	"log"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego/orm"
)

func init() {

	orm.RegisterDataBase("default", "mysql", beego.AppConfig.DefaultString("mysql", "root:123456@tcp(127.0.0.1:3306)/rworld?charset=utf8"), 30)

	// register model
	orm.RegisterModel(new(Task))
	orm.RegisterModel(new(RTest))
	orm.RegisterModel(new(JenkinsNode))
	orm.RegisterModel(new(Product))
	orm.RegisterModel(new(ProductVersion))

	orm.RegisterModel(new(AutoTask))

	// create table
	orm.RunSyncdb("default", false, true)

	DefaultJenkinsNodeList = NewJenkinsNodeManager()
	DefaultJenkinsNodeList.nodes = DefaultJenkinsNodeList.SelectAll()
	DefaultJenkinsNodeList.ConnectAll()

	DefaultProductList = NewProductManager()
	DefaultProductList.Products = DefaultProductList.SelectAll()

	for _, p := range DefaultProductList.Products {
		log.Printf("-------Product%v= %v,jenkins_node=%v\n", p.Id, p, p.JenkinsNode)
	}
	DefaultProductVersionManageList = NewProductVersionManage()
	DefaultProductVersionManageList.Versions = DefaultProductVersionManageList.SelectAll()
	for _, v := range DefaultProductVersionManageList.Versions {
		log.Printf("-------Version%v = %v  ,Product=\n", v.Id, v, v.Product)
	}
	DefaultAutoTaskList = NewAutoTaskManager()
}
