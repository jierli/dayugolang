package main

import (
	"fmt"
	"time"
)

func chars(prefix string) {
	for c := 'A'; c <= 'Z'; c++ {
		fmt.Printf("%s: %c\n", prefix, c)
		time.Sleep(time.Microsecond * 3)
	}
}

func main() {
	//go 调用实现并发编程
	go chars("goruting")
	chars("main")
	//time.Sleep(time.Second * 3)
}
