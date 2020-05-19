package main

import (
	"fmt"
	"strings"
)

func main() {
	//定义builder 结构体对象，类似定义内存中的流对象（写文件对象）
	var builder strings.Builder
	//builder写入内存字节切片
	builder.Write([]byte("dayu test\n"))
	//builder写入内存字符串
	builder.WriteString("字符串 dayutest\n")
	fmt.Println(builder.String())
}
