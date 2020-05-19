package main

import (
	"fmt"
	"time"
)

type User struct {
	id   int
	name string
	addr string
	tel  string
}

//自定义结构体类型
type Task struct {
	id        int
	name      string
	startTime time.Time
	endTime   time.Time
	status    int
	User      //匿名嵌入--》类似面向对象的继承，也有一个属性名User（默认简写）
}

func main() {
	var task Task
	fmt.Printf("%#v\n", task)

	//访问user 的执行人名
	fmt.Println(task.User.name)
	//修改
	task.User.name = "dayu_anoy_v1"
	fmt.Println(task)

	//赋值
	task = Task{
		id:   1,
		name: "完成tudolist",
		User: User{
			id:   1,
			name: "dayu_anoy_v2",
			addr: "北京",
		},
	}
	fmt.Printf("%#v\n", task)
	/*
	匿名嵌入与命名嵌入的区别。访问时命名嵌入必须指定结构体实例名（也就是完整的访问关系）
	当匿名嵌入的多个对象有相同的属性名，必须通过全名访问
	*/
	fmt.Println(task.name, task.addr) 。
	task.addr = "新加坡"
	fmt.Println(task.name, task.addr)
}
