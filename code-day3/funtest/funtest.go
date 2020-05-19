package main

import "fmt"

// 声明一个函数类型
type cb func(int, int) int

func main() {
	testCallBack(1, 111, callBack)
	testCallBack(2, 222, func(x int, y int) int {
		fmt.Printf("我是回调，x：%d\n", x)
		return x
	})
}

func testCallBack(x int, y int, f cb) {
	f(x, y)
	fmt.Printf("functestCallBack:%#v\n", y)

}

func callBack(x int, y int) int {
	fmt.Printf("我是回调，x：%d\n", x)
	fmt.Printf("func_callBack:%#v\n", y)
	return x
}
