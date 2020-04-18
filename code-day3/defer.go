package main

import "fmt"

func main() {
	//defer 函数调用--在函数退出前执行
	defer func() {
		fmt.Println("defer 函数调用1")
	}()
	defer func() {
		fmt.Println("defer 函数调用2")
	}()
	defer func() {
		fmt.Println("defer 函数调用3")
	}()
	fmt.Println("main函数调用")
}
