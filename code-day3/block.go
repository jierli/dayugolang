package main

import "fmt"

func main() {
	name, desc := "dayu", "kao kk"
	func(name string) {
		fmt.Println(name, desc)
		name, desc = "dayu10000000", "kao kk1000000000"
		//name, desc := "dayu10000000", "kao kk1000000000"  注意短申明作用域到块
		fmt.Println(name, desc)
	}("test")
	fmt.Println(name, desc)
}
