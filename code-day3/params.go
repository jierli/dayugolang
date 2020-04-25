package main

import "fmt"

//参数类型合并
//连续多个变量类型相同
//保留最后一个元素类型，前面的类型都可以省略

func add(n1, n2 int) int {
	return n1 + n2
}

//可变数量传入参数
func test(args ...string) {
	fmt.Printf("%#v\n", args)
}

func main() {
	a := add(1, 2)
	fmt.Println(a)
	test("a", "b")
	test3(1, 2, 3, 4, 5, 6)
	test3(11, 21, 31, 41, 51, 61)
	par := []int{11, 12, 13, 14, 15}
	test3(1, 2, par...) // ...解切片操作
}

//可变参数中要保留最少数量传参
func test3(n1, n2 int, args ...int) {
	togal := n1 + n2
	for _, var1 := range args {
		togal += var1
		fmt.Println(var1)
	}
}

//可变参数传入给另一个函数的可变参数
func calc(n1, n2 int, args ...int) {
	//调用test3 传入args 可变参数
	test3(n1, n2, args...) //解切片操作,展开可变参数，传递给另外一个函数
}
