package main

import "fmt"

func test1(n int) {
	fmt.Printf("%v:%d\n", &n, n)
	n = 1

}

func test2(s []int) {
	fmt.Printf("%p\n", s)
	s[0] = 1
}

func main() {
	a := test1 //把函数赋值给a

	fmt.Printf("%T\n", a)
	a(2222)
	var callback func(int) //定义函数类型
	fmt.Printf("%T,%#v\n", callback, callback)

	callback = test1 //函数赋值给函数类型变量
	fmt.Printf("%T,%#v\n", callback, callback)
	callback(10000)
}

//函数变量的作用，可以吧函数当做传参，传入给另一个函数f
