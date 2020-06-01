package main

import (
	"fmt"
)

func main() {
	//定义管道，元素为int
	channel := make(chan int)

	for i := 0; i < 3; i++ {
		go func(prefix int) {
			for ch := 'A'; ch <= 'Z'; ch++ {
				fmt.Printf("%d,%c\n", prefix, ch)
			}
			channel <- prefix
		}(i)

		fmt.Printf("%#v\n", channel)
	}
	//close(channel)
	for i := 0; i < 3; i++ {
		num := <-channel
		fmt.Println(num)
	}

}
