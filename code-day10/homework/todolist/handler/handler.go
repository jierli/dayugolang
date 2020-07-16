package handler

import (
	"fmt"
	"time"
	"html/template"
	"net/http"
)

type Task struct {
	ID int
	Name string
	Status int
	Content string
	StartTime time.Time
	Deadline time.Time
	CompleteTime *time.Time
	User string
}
/*
type TaskForm struct {
	ID int
	Name string
	Status int
	Content string
	StartTime time.Time
	Deadline time.Time
	CompleteTime *time.Time
	User string
}
*/
func (h *Task) Index(response http.ResponseWriter, request *http.Request)  {
	tasks:=make([]Task,0,20)
	task:=Task{
		ID:1,
		Name :"tudolist",
		Status:1,
		Content:"test",
		User :"dayu",
	}
	tasks=append(tasks,task)
	fmt.Println(tasks)


	tpl := template.Must(template.ParseFiles("views/index.html"))
	tpl.ExecuteTemplate(response, "index.html", struct {
		Tasks []Task
	}{tasks})
	fmt.Println(tasks)
}
