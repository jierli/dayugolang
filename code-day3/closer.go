package main

import "fmt"

//函数可以当做返回值
//变量的生命周期，一般是方法执行结束后结束，语言会做垃圾回收
//闭包
func addBase(base int) func(int) int {
	fmt.Println(base)
	return func(n int) int {
		return n + base
	}
}

func main() {
	add1 := addBase(11)
	fmt.Println(add1(5))

	add100 := addBase(100)
	fmt.Println(add100(5))
}
