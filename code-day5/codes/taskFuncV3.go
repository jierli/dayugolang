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

func NewUser(id int, name, addr, tel string) *User {
	return &User{id, name, addr, tel}
}

//结构体方法
func (u *User) SetName(name string) {
	u.name = name
}
func (u *User) SetAddr(addr string) {
	u.addr = addr
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

func NewTask(id int, name string, startTime, endTime *time.Time, status int, user *User) *Task {
	return &user{id, name, startTime, endTime, status, user}
}

func main() {
	now := time.Now()
	user := NewUser(1, "dayu", "", "")
	task := NewTask(1, "完成知识整理", &now, &now, user)

	task.User.SetName("大雨")
	task.User.SetAddr("厦门")

	fmt.Printf("%#v", task)
	fmt.Printf("%#v", task.User)
}
