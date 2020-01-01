package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

var DefaultTaskList *TaskManager

type Task struct {
	Id      int64  // Unique identifier
	Title   string // Description
	Done    bool   // Is this task done?
	Des     string // Is this task done?
	Bgntime time.Time
	Endtime time.Time
}

// NewTask creates a new task given a title, that can't be empty.
func NewTask(title string) (*Task, error) {
	if title == "" {
		return nil, fmt.Errorf("empty title")
	}
	o := orm.NewOrm()
	t := Task{Title: title}
	num, err := o.InsertOrUpdate(&t)

	if err == nil && num > 0 {
		return &t, nil
	}
	return nil, err
}

// TaskManager manages a list of tasks in memory.
type TaskManager struct {
	tasks  []*Task
	lastID int64
}

// NewTaskManager returns an empty TaskManager.
func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

// Save saves the given Task in the TaskManager.
func (m *TaskManager) Insert(task *Task) error {
	o := orm.NewOrm()
	o.Begin()
	num, err := o.Insert(task)

	if err == nil && num > 0 {
		o.Commit()
		return nil
	}
	o.Rollback()
	return err
}

// Find returns the Task with the given id in the TaskManager and a boolean
// indicating if the id was found.
func (m *TaskManager) Update(t Task) (maps []orm.Params, err error) {

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

// Save saves the given Task in the TaskManager.
func (m *TaskManager) Save(task *Task) error {
	o := orm.NewOrm()
	num, err := o.InsertOrUpdate(task)

	if err == nil && num > 0 {
		return nil
	}
	return err
}

// cloneTask creates and returns a deep copy of the given Task.
func cloneTask(t *Task) *Task {
	c := *t
	return &c
}

func (m *TaskManager) SelectAll() []*Task {
	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw("select * FROM task; ").Values(&maps)

	if err != nil && num > 0 {
		return nil
	}
	var tasks []*Task
	for _, mm := range maps {
		tmpId, _ := strconv.ParseInt(mm["id"].(string), 10, 64)
		tmpDone, _ := strconv.ParseBool(mm["done"].(string))
		bgnTime, _ := time.Parse("2006-01-02 15:04:05", mm["bgntime"].(string))
		endTime, _ := time.Parse("2006-01-02 15:04:05", mm["endtime"].(string))
		tasks = append(tasks, &Task{Id: tmpId, Title: mm["title"].(string), Done: tmpDone,
			Des:     mm["des"].(string),
			Bgntime: bgnTime,
			Endtime: endTime})
	}
	fmt.Println(tasks)
	return tasks

}

// Find returns the Task with the given id in the TaskManager and a boolean
// indicating if the id was found.
func (m *TaskManager) Find(ID int64) (maps []orm.Params, err error) {

	o := orm.NewOrm()
	num, err := o.Raw("select * FROM task WHERE ID=? ", ID).Values(&maps)

	if err == nil && num > 0 {

		return maps, nil
	}
	return nil, err
}

func init() {
	DefaultTaskList = NewTaskManager()
}
