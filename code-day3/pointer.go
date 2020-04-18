package main

import "fmt"

//值类型，在函数内修改实参的值
func chage(value int) {
	value = value + 1
}
func chagePoint(poinert *int) {
	*poinert = *poinert + 1
}
func main() {
	value := 1
	chage(value)
	fmt.Println(value)
	chagePoint(&value)
	fmt.Println(value)
}
