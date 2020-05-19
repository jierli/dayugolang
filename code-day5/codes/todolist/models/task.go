package models

import "time"

type User struct {
	id   int
	name string
	addr string
	tel  string
}

//自定义结构体类型
type taskv1 struct {
	id        int
	name      string
	startTime time.Time
	endTime   time.Time
	status    int
	user      *User //匿名嵌入--》类似面向对象的继承，也有一个属性名User（默认简写）
}

//自定义结构体类型
type Taskv2 struct {
	id        int
	name      string
	startTime time.Time
	endTime   time.Time
	status    int
	user      *User //匿名嵌入--》类似面向对象的继承，也有一个属性名User（默认简写）
}

type Taskv3 struct {
	Id        int
	Name      string
	StartTime time.Time
	EndTime   time.Time
	Status    int
	User      *User //匿名嵌入--》类似面向对象的继承，也有一个属性名User（默认简写）
}

type anonyWrapperv1 struct {
	taskv1
}

type AnonyWrapperv2 struct {
	Taskv2
}

type AnonyWrapperv3 struct {
	Taskv3
}
