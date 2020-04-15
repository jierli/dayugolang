package main

import "fmt"


//bread + goto 到指定的标记位置
func main() {
BREAK:
	for j := 0; j < 3; j++ {
		fmt.Println(j, "------")
		for i := 0; i < 5; i++ {
			fmt.Println(i)
			if i == 3 {
				break BREAK
			}
		}
	}

}
