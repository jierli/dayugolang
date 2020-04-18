package main

import "fmt"

//匿名函数
func calc(n1 int, n2 int, callback func(int, int) int) int {
	//不定义什么运算
	//通过函数参数传递给我要进行的运算
	rt := callback(n1, n2)
	//检查结果在0到100之间
	if rt >= 0 && rt <= 100 {
		return rt
	}
	return -1
}

func main() {
	//临时使用函数
	//函数直接赋值给变量
	mult := func(n1, n2 int) int {
		return n1 * n2
	}
	add := func(n1, n2 int) int {
		return n1 + n2
	}

	rt := calc(20, 50, func(n1, n2 int) int {
		return n1 * n2
	})
	fmt.Println("匿名*法", rt)
	rt = calc(20, 50, func(n1, n2 int) int {
		return n1 + n2
	})
	fmt.Println("匿名+法", rt)

	rt = calc(20, 50, mult)
	fmt.Println(rt)
	rt = calc(20, 50, add)
	fmt.Println(rt)

	//方法调用
	/*
		类似
		a:=func(int)
		a()
	*/
	func() {
		fmt.Println("方法调用")
	}()
	func(i int) {
		fmt.Println(i)
	}(100)
}
