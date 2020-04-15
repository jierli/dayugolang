package main

import "fmt"

func main()  {
	//nil切片与空切片区别
	var nilSlice []int
	var emptySlice []int=[]int{}

	fmt.Printf("%T,%#v\n",nilSlice,nilSlice)
	fmt.Printf("%T,%#v\n",emptySlice,emptySlice)
	nilSlice=append(nilSlice,1)
	emptySlice=append(emptySlice,1)
	fmt.Printf("%T,%#v\n",nilSlice,nilSlice)
	fmt.Printf("%T,%#v\n",emptySlice,emptySlice)

}