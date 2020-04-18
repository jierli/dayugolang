package main

import "fmt"

func main() {
	//切片操作的应用场景
	//队列，先进先出
	queue := []string{}
	//push
	//append
	queue = append(queue, "a", "b")
	queue = append(queue, "c")
	fmt.Println(queue)
	//pop 从左边删元素
	x := queue[0]
	queue = queue[1:]
	fmt.Println("1:", x)
	x = queue[0]
	queue = queue[1:]
	fmt.Println("2:", x)
	x = queue[0]
	queue = queue[1:]
	fmt.Println("3:", x)

	//栈 从左边删元素
	queue = append(queue, "a", "b")
	queue = append(queue, "c")
	fmt.Println(queue)
	x = queue[len(queue)-1]
	queue = queue[:len(queue)-1]
	fmt.Println("daoxu1:", x)
	x = queue[len(queue)-1]
	queue = queue[:len(queue)-1]
	fmt.Println("daoxu1:", x)
	x = queue[len(queue)-1]
	queue = queue[:len(queue)-1]
	fmt.Println("daoxu1:", x)

}
