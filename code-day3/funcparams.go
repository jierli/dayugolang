package main

import "fmt"

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

func mult(n1, n2 int) int {
	return n1 * n2
}
func add(n1, n2 int) int {
	return n1 + n2
}

func main() {
	result := calc(11, 12, mult)
	fmt.Println(result)
	result = calc(11, 12, add)
	fmt.Println(result)
}
