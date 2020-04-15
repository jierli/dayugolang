package main

import (
	"fmt"
	"sort"
)

//切片数组切片包
func main() {
	//排序
	numbs := []int{2, 5, 1, 27, 82, 4}
	fmt.Println("排序前", numbs)
	sort.Ints(numbs)
	fmt.Println("排序后", numbs)

	names := []string{"a", "u", "d", "n"}
	fmt.Println("排序前", names)
	sort.Strings(names)
	fmt.Println("排序后", names)

	//查找
}
