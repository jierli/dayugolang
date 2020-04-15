package main

import "fmt"

//冒泡排序算法
func main() {
	arr := []int{6, 1, 4, 7, 5, 3, 2}
	for j := 0; j < len(arr)-1; j++ {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
			fmt.Println(i, arr)
		}
	}

}
