package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

var DefaultAutoTaskList *AutoTaskManager

type AutoTask struct {
	Id             int64  // Unique identifier
	Name           string //
	Des            string //
	SubmitFilePath string
	ProductVersion *ProductVersion `orm:"rel(fk)"`
}

func NewAutoTask(name string) (*AutoTask, error) {
	if name == "" {
		return nil, fmt.Errorf("empty name")
	}

	o := orm.NewOrm()
	t := AutoTask{Name: name}
	num, err := o.InsertOrUpdate(&t)

	if err == nil && num > 0 {
		return &t, nil
	}
	return nil, err
}

type AutoTaskManager struct {
	autotasks []*AutoTask
	lastID    int64
}

func NewAutoTaskManager() *AutoTaskManager {
	return &AutoTaskManager{}
}

func (m *AutoTaskManager) Insert(t *AutoTask) error {
	o := orm.NewOrm()
	o.Begin()
	num, err := o.Insert(t)

	if err == nil && num > 0 {
		o.Commit()
		return nil
	}
	o.Rollback()
	return err
}

func (m *AutoTaskManager) Update(t AutoTask) (maps []orm.Params, err error) {

	o := orm.NewOrm()
	o.Begin()
	num, err := o.Update(&t)

	if err == nil && num > 0 {
		o.Commit()
		return maps, nil
	}
	o.Rollback()
	return nil, err
}

func (m *AutoTaskManager) Save(t *AutoTask) error {
	o := orm.NewOrm()
	num, err := o.InsertOrUpdate(t)

	if err == nil && num > 0 {
		return nil
	}
	return err
}

func cloneAutoTask(t *AutoTask) *AutoTask {
	c := *t
	return &c
}

func (m *AutoTaskManager) SelectAll() []*AutoTask {
	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw("select * FROM autotask; ").Values(&maps)

	if err != nil && num > 0 {
		return nil
	}
	var autoTasks []*AutoTask
	for _, mm := range maps {
		tmpId, _ := strconv.ParseInt(mm["id"].(string), 10, 64)
		//tmpJenkinsNodeId, _ := strconv.ParseInt(mm["jenkins_node_id"].(string), 10, 64)

		autoTasks = append(autoTasks, &AutoTask{Id: tmpId, Name: mm["name"].(string),
			Des: mm["des"].(string)})
	}
	fmt.Println(autoTasks)
	return autoTasks

}

func (m *AutoTaskManager) Find(ID int64) (maps []orm.Params, err error) {

	o := orm.NewOrm()
	num, err := o.Raw("select * FROM autotask WHERE ID=? ", ID).Values(&maps)

	if err == nil && num > 0 {

		return maps, nil
	}
	return nil, err
}
