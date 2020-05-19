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
	//方法
	//方法值 实例.方法名
	task := Task{}
	task2 := &Task{}

	methodValue1 := task.SetName
	methodValue2 := task2.SetName //自动取引用

	fmt.Printf("%#v\n", methodValue1)
	fmt.Printf("%#v\n", methodValue2)
	//方法表达式  结构体.方法名
	methodValue1("完成作业")
	methodValue2("完成作业")
	fmt.Printf("%#v\n", task)
	fmt.Printf("%#v\n", task2)

	methodValue3 := task.SetUser
	methodValue4 := task2.SetUser //自动取引用

	fmt.Printf("%#v\n", methodValue3)
	fmt.Printf("%#v\n", methodValue4)
	//方法表达式  结构体.方法名
	methodValue3("大雨")
	methodValue4("大雨")
	fmt.Printf("%#v\n", task)
	fmt.Printf("%#v\n", task2)

}
