package main

import "fmt"

func main() {
	//移除切片操作
	//移除第一个元素或者移除最后一个元素
	numbs := []int{0, 1, 3, 4, 5, 6, 7, 8, 9}
	numbs = numbs[1:]
	fmt.Println(numbs)
	//移除最后一个元素
	numbs = []int{0, 1, 3, 4, 5, 6, 7, 8, 9}
	numbs = numbs[:len(numbs)-1]
	fmt.Println(numbs)

	//利用切片与数组共享内存空间原理，与copy 操作实现删除中间元素
	//删除元素3的实现
	numbs = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	copy(numbs[3:], numbs[4:])
	fmt.Printf("%v\n", numbs)
	numbs = numbs[:len(numbs)-1]
	fmt.Printf("%v\n", numbs)

}
