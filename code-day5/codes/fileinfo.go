package main

import (
	"fmt"
	"os"
)

//获取文件信息
func main() {
	fileInfo, err := os.Stat("append.log")
	fmt.Println(fileInfo, err)
}
