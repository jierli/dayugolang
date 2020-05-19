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
	user      User //结构体 命名嵌入，在面向对象里面叫组合
}

func main() {
	var task Task
	fmt.Printf("%#v\n", task)

	//访问user 的执行人名
	fmt.Println(task.user.name)
	//修改
	task.user.name = "dayu_v1"
	fmt.Println(task)

	//赋值
	task = Task{
		id:   1,
		name: "完成tudolist",
		user: User{
			id:   1,
			name: "dayu_user_v2",
		},
	}
	fmt.Printf("%#v\n", task)
}
