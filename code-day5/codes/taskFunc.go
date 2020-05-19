package main

import (
	"fmt"
	"time"
)

type Task struct {
	id        int
	name      string
	startTime *time.Time
	etartTime *time.Time
	user      string
}
//结构体方法
//定义方法
//特定类型指定--接收者
//接收这（类型名   类型）
func (task *Task) SetName(name string) {
	task.name = name
}
func (task *Task) GetName() string {
	return task.name
}

func main() {
	task := &Task{name: "完成TODO"}
	task.SetName("知识整理")

	fmt.Println(task.GetName())
}
