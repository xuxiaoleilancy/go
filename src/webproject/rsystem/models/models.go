package models

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego/orm"
)

func init() {

	// set default database
	//orm.RegisterDataBase("default", "mysql", "root:123456@tcp(10.10.8.28)/rworld?charset=utf8", 30)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/rworld?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(Task))
	orm.RegisterModel(new(RTest))

	// create table
	orm.RunSyncdb("default", false, true)

	DefaultTaskList = NewTaskManager()
	DefaultRTestList = NewRTestManager()
}
