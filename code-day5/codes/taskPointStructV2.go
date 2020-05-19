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
	fmt.Printf("%#v", task)

	task = Task{
		id:   1,
		name: "完成tudolist",
		User: &User{
			id:   1,
			name: "dayu_porint_v3",
			addr: "北京",
		},
	}
	fmt.Printf("%#v\n", task)
	fmt.Println(task.name)
	fmt.Println(task.User.name)
	fmt.Println(task.User.addr)
}
