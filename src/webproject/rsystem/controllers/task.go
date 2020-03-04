package controllers

import (
	"encoding/json"

	"strconv"
	"time"
	"webproject/rsystem/models"

	"github.com/astaxie/beego"
)

type TaskController struct {
	beego.Controller
}

// Example:
//
//   req: GET /todo/task/
//   res: 200 {"Tasks": [
//          {"ID": 1, "Title": "Learn Go", "Done": false},
//          {"ID": 2, "Title": "Buy bread", "Done": true}
//        ]}
func (this *TaskController) ListTasks() {

	maps := models.DefaultTaskList.SelectAll()

	res := struct{ Tasks []*models.Task }{maps}
	this.Data["json"] = res
	this.ServeJSON()

	defer this.Render()
}

// Examples:
//
//   req: POST /task/ {"Title": ""}
//   res: 400 empty title
//
//   req: POST /task/ {"Title": "Buy bread"}
//   res: 200
func (this *TaskController) NewTask() {

	req := struct{ Title string }{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("empty title"))
		return
	}
	t := models.Task{Title: req.Title, Bgntime: time.Now(), Endtime: time.Now().Add(time.Duration(time.Hour * 24))}
	models.DefaultTaskList.Insert(&t)

	this.Render()
}

// Examples:
//
//   req: GET /task/1
//   res: 200 {"ID": 1, "Title": "Buy bread", "Done": true}
//
//   req: GET /task/42
//   res: 404 task not found
func (this *TaskController) GetTask() {
	id := this.Ctx.Input.Param(":id")
	beego.Info("Task is ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	t, err := models.DefaultTaskList.Find(intid)

	defer this.Render()
	if err != nil {

		beego.Info("Found", err)
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("task not found"))

	} else {

		this.Data["json"] = t
		this.ServeJSON()
	}
}

// Example:
//
//   req: PUT /task/1 {"ID": 1, "Title": "Learn Go", "Done": true}
//   res: 200
//
//   req: PUT /task/2 {"ID": 2, "Title": "Learn Go", "Done": true}
//   res: 400 inconsistent task IDs
func (this *TaskController) UpdateTask() {
	id := this.Ctx.Input.Param(":id")
	beego.Info("Task is ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	var t models.Task
	defer this.Render()
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &t); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	if t.Id != intid {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("inconsistent task IDs"))
		return
	}
	models.DefaultTaskList.Update(t)
}
