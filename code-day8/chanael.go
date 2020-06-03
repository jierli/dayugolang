package main

import (
	"fmt"
	"time"
)

func main() {
	//定义管道，元素为int
	var channel chan int

	fmt.Printf("%T,%v,%d\n", channel, channel, len(channel))

	//初始化（无缓冲区的管道）
	channel = make(chan int, 3)
	fmt.Printf("%T,%v,%d\n", channel, channel, len(channel))

	//管道是缓冲区
	//写  <-  channel <- 1  无缓冲区管道  必须有另外一个例程进行读取,在不能写入数据时会进行阻塞。
	go func() {
		time.Sleep(time.Second * 5)
		channel <- 1
		channel <- 1
		channel <- 1
	}()

	//读   <- channel 无缓冲区管道  必须有另外一个例程进行写取,在没有读到数据会进行阻塞。
	fmt.Println(len(channel))
	num := <-channel
	fmt.Println(num, len(channel))
	num = <-channel
	fmt.Println(num, len(channel))
	num = <-channel
	fmt.Println(num, len(channel))

}
