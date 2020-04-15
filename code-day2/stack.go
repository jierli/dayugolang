package main

import "fmt"

func main() {
	//堆栈
	//先进后出，电梯原理，先上的人后出，后上的人先出。
	stack := []string{}
	//push
	//append
	stack = append(stack, "a")
	stack = append(stack, "b")
	stack = append(stack, "c")
	fmt.Println(stack)
	//pop
	//后面移除
	x := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println("发射1", x)
	x = stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println("发射2", x)
	x = stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println("发射3", x)
}
