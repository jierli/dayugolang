package main

import (
	"fmt"
	"strings"
)

func main() {
	//通过new函数创建reader结构体指针，初始化字符串
	reader := strings.NewReader("abc12345678xyz")

	//定义切片用于存放读取的内容
	ctx := make([]byte, 10)
	for {
		i := 1
		//读取内容到切片，n读取字节数量
		n, err := reader.Read(ctx)
		if err != nil {
			break
		}
		fmt.Println(i, n, err, string(ctx[:n]))
		i++
	}

}
