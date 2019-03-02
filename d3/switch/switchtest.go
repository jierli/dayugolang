package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	m, err1 := strconv.Atoi(os.Args[1])
	n, err2 := strconv.Atoi(os.Args[3])
	fmt.Println(os.Args[2])
	if err1 == nil && err2 == nil {
		switch os.Args[2] {
		case "+":
			fmt.Println(m + n)
		case "-":
			fmt.Println(m - n)
		case "*":
			fmt.Println(m * n)
		case "/":
			fmt.Println(m / n)

		}
	}
}
