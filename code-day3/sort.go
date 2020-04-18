package main

import (
	"fmt"
	"sort"
)

func main() {

	stats := [][]int{{'A', 2}, {'B', 4}, {'C', 5}, {'D', 2}}
	fmt.Printf("%v\n", stats)
	//使用次数排序
	//稳定排序，不稳定排序
	sort.Slice(stats, func(i, j int) bool { return stats[i][1] < stats[j][1] })
	fmt.Printf("%v\n", stats)
}
