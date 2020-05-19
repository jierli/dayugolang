package main

import "fmt"

//自定义类型
type Counter int
type Task map[string]string
type Collback func(...string)

func main() {
	var cnt Counter
	fmt.Printf("%T\n", cnt)
	fmt.Printf("%#v\n", cnt)

	cnt = 1
	fmt.Printf("%T\n", cnt)
	fmt.Printf("%#v\n", cnt)

	var task Task
	fmt.Printf("%T\n", task)
	fmt.Printf("%#v\n", task)
	task = map[string]string{"name": "dayu"}
	fmt.Printf("%T\n", task)
	fmt.Printf("%#v\n", task)

	//自定义函数类型
	var print Collback
	print = func(args ...string) {
		for i, e := range args {
			fmt.Printf("%#v,%#v", i, e)
		}
	}
	print("a", "b", "c")
}
