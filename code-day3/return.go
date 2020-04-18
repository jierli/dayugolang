package main

import "fmt"

func mult(n1, n2 int) int {
	return n1 * n2
}
func add(n1, n2 int) int {
	return n1 + n2
}
func calc(n1, n2 int) (int, int) {
	r1 := mult(n1, n2)
	r2 := add(n1, n2)
	return r1, r2
}

//命名返回值
func calc2(n1, n2 int) (r1 int, r2 int) {
	r1 = mult(n1, n2)
	r2 = add(n1, n2)
	return //不需要指定变量
}

func main() {

	fmt.Println(calc(1, 2))
	fmt.Println(calc2(1, 2))
}
