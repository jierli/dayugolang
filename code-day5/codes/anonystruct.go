package main

import "fmt"

func main() {
	//匿名结构体--》使用方法：直接初始化给一个变量
	//只能在函数内使用
	user := struct {
		id   int
		name string
		age  int
	}{1, "dayu", 34}
	fmt.Printf("%T,%#v\n", user, user)
	fmt.Println(user.name)
	user.name = "dayu anonystruct"
	fmt.Println(user)

	//你们结构体字面量初始化
	user = struct {
		id   int
		name string
		age  int
	}{1, "dayu", 34}
}
