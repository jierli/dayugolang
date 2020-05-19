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
	task := Task{}   //值
	task2 := &Task{} //指针
	fmt.Println(task, task2)
	//方法表达式  结构体.方法名
	//对于值接收者，可以通过指针/值来获取
	method1 := Task.SetName
	method2 := (*Task).SetName //自动取引用  //函数类型func(*main.Task, string)

	//method3 := Task.SetName
	//对于指针接收者，只可以通过指针
	method4 := (*Task).SetUser //自动取引用

	fmt.Printf("%T\n", method1)
	fmt.Printf("%T\n", method2)
	fmt.Printf("%T\n", method4)

}
