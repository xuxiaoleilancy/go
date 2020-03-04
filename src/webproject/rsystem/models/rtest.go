package models

import (
	"strconv"

	"github.com/astaxie/beego/orm"
)

var DefaultRTestList *RTestManager

type RTest struct {
	Id     int64  // Unique identifier
	Title  string // Title
	Des    string // Des?
	Params string
	Help   string
}

// TaskManager manages a list of tasks in memory.
type RTestManager struct {
	cmds []*RTest
}

// NewTaskManager returns an empty TaskManager.
func NewRTestManager() *RTestManager {
	return &RTestManager{}
}

func (m *RTestManager) SelectAll() []*RTest {
	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw("select * FROM r_test; ").Values(&maps)

	if err != nil && num > 0 {
		return nil
	}
	var cmds []*RTest
	for _, mm := range maps {
		tmpId, _ := strconv.ParseInt(mm["id"].(string), 10, 64)
		cmds = append(cmds, &RTest{Id: tmpId, Title: mm["title"].(string), Params: mm["params"].(string),
			Des:  mm["des"].(string),
			Help: mm["help"].(string)})
	}
	return cmds

}
