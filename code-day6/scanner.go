package main

import (
	"bufio"
	"fmt"
	"os"
)

//带缓冲的io
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("请输入内容")
		scanner.Scan()
		if "q" == scanner.Text() {
			break
		}
		fmt.Println("输入内容是:", scanner.Text())
	}

}
