package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var (
		x  byte
		x1 int
		x2 int32
		x3 int64
		x4 uint
		x5 uint32
		x6 uint64
		x7 int8
		x8 uint8
	)
	x7 = 127
	x7 = x7 + 1
	x8 = 20
	fmt.Println(x, x1, x2, x3, x4, x5, x6)
	fmt.Println(unsafe.Sizeof(x1))
	fmt.Println(unsafe.Sizeof(x3))
	fmt.Println(x7, x8)
	str1 := "hello"
	doc := ``
	fmt.Println(str1,doc)
}
