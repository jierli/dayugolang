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
	*User     //匿名嵌入--》类似面向对象的继承，也有一个属性名User（默认简写）
}

func main() {
	var task Task
	fmt.Printf("%#v\n", task)

	task = Task{
		id:   1,
		name: "完成tudolist",
		User: &User{
			id:   1,
			name: "dayu_porint_v3",
			addr: "北京",
		},
	}

	task2 := task
	task2.name = "完成整理"

	fmt.Println(task)
	fmt.Println(task2)

	task2.User.name = "dayu_point_v4"
	fmt.Println(task.User)
	fmt.Println(task2.User)
}
