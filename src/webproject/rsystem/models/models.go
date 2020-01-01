package models

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego/orm"
)

func init() {

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/rworld?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(Task))

	// create table
	orm.RunSyncdb("default", false, true)

	DefaultTaskList = NewTaskManager()
}
