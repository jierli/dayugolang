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

//定义方法
//特定类型指定--接收者
//接收这（类型名   类型）
// 读用值接受类型，写用指针接收者
func (task Task) SetName(name string) {
	task.name = name
}

//user 指针接收者
func (task *Task) SetUser(user string) {
	task.user = user
}

func main() {
	task := Task{name: "完成TODO"}
	task.SetName("知识整理")

	//go 的语法糖 go编译时自动转换(&task).SetUser("大雨point")  称之为取引用
	task.SetUser("大雨point") //类似(&task).SetUser("大雨point")

	fmt.Printf("%#v\n", task)

	task2 := &Task{name: "完成TODO"} //解引用
	task2.SetName("知识整理")
	task2.SetUser("大雨point_v2")
	fmt.Printf("%#v\n", task2)
}
