package main

import (
	"fmt"
	"time"
)

//自定义结构体类型
type Task struct {
	id        int
	name      string
	startTime time.Time
	endTime   time.Time
	status    int
	user      string
}

func main() {
	//使用结构体类型
	var task Task //结构体实例
	fmt.Printf("%T\n", task)
	fmt.Printf("%#v\n", task)

	//结构体零值
	task = Task{}
	fmt.Printf("%#v\n", task)

	//结构体赋值
	//第一种赋值：按照属性顺序创建Task字面量进行赋值
	task = Task{1, "完成Todolist", time.Now(), time.Now().Add(24 * time.Hour), 1, "dayu"}
	fmt.Printf("%T\n", task)
	fmt.Printf("%#v\n", task)

	//第二种赋值：按照属性名进行赋值，可部分赋值
	//换行赋值最后一个逗号不能省略
	task = Task{id: 2,
		name: "完成知识整理",
		user: "dayu",
	}
	fmt.Printf("%T\n", task)
	fmt.Printf("%#v\n", task)

	//结构体指针类型实例
	var ptask *Task
	fmt.Printf("%T\n", ptask)
	fmt.Printf("%#v\n", ptask)
	//指针结构体类型赋值
	ptask = &Task{}
	fmt.Printf("%#v\n", ptask)

	//结构体实例：访问属性，修改属性
	fmt.Println(task.id, task.name, task.user)
	//修改
	task.user = "dayu_update"
	fmt.Println(task.id, task.name, task.user)

	//结构体实例 值类型，引用类型判断
	//判断结果是值类型
	task2 := task
	task2.user = "dayu_update_task2"
	fmt.Println(task.id, task.name, task.user)
	fmt.Println(task2.id, task2.name, task2.user)

	//new函数初始化结构体
	//经测试为引用类型
	task3 := new(Task)
	fmt.Printf("%T\n", task3)
	fmt.Printf("%#v\n", task3)

	task3 = &Task{id: 3,
		name: "完成知识整理_new",
		user: "dayu_new",
	}
	fmt.Printf("%#v\n", task3)
	task4 := task3
	task4.user = "dayu_update_task4"
	fmt.Println(task3.id, task3.name, task3.user)
	fmt.Println(task4.id, task4.name, task4.user)
}
