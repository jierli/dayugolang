package main

import (
	"fmt"
)

func main() {
	//定义管道，元素为int
	channel := make(chan struct{})

	go func() {
		channel <- struct{}{}
		close(channel)

	}()

	//通过ok判断是不是关闭管道
	num, ok := <-channel
	fmt.Println(num, ok)
}
