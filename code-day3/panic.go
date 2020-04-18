package main

import "fmt"

func test() (err error) {
	defer func() {
		fmt.Println("defer")
		if panicerr := recover(); panicerr != nil { //recover错误捕获
			fmt.Printf("%T,%v\n", panicerr, panicerr) //recover 必须在延迟函数执行
			err = fmt.Errorf("%s", panicerr)
		}

	}()
	fmt.Println("before")
	panic("自定义panic") //主动显示抛出错误，默认情况下抛错后退出执行。加入recover 机制捕获错误，就能继续执行。
	fmt.Println("after")
	return
}

func main() {
	fmt.Println("main before")
	err := test()
	fmt.Println("main after")
	fmt.Println(err)
}
