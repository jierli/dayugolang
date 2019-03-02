package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	i := 5
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	//遍历数组元素
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
	s := "dayuTEST遍历数组元素"
	for i, str := range s {
		fmt.Println(i, string(str))
	}

	a := 5
	for {
		a = a + 1
		fmt.Println(a)
		if a > 7 {
			break
		}
	}
}
