package main

import "fmt"

func main() {
	ascii := "abcd我爱中华人民共和国"

	fmt.Println([]byte(ascii))
	fmt.Println([]rune(ascii))
	//统计字符的数量
	fmt.Println(len(ascii))
	fmt.Println(len([]rune(ascii)))
}
