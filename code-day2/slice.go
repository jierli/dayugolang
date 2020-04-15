package main

import "fmt"

/*
切片长度可变，底层是通过数组实现，当增加切片时按照一定的增长因子重新申请内存空间。
所以长度和容量时不一样的
len()获取长度 cap()获取容量
*/
func main() {
	var names []string
	//零值  nil
	fmt.Printf("%#v\n", names)

	//初始化
	//字面量
	names = []string{} //空切片 空值
	fmt.Printf("%#v\n", names)
	//切片长度可变
	//赋值1 直接赋值
	names = []string{"test1", "test2", "test3"}
	fmt.Printf("%#v\n", names)
	//赋值2 通过索引赋值
	names = []string{1: "dayu1", 10: "dayu10"}
	fmt.Printf("%#v\n", names)

	//赋值3 make函数
	//make方法1 两参数只指定类型与长度
	names = make([]string, 5)
	fmt.Printf("两参数只指定类型与长度%#v\n", names)
	//make方法2 三参数指定类型与长度与容量
	names = make([]string, 3, 10)
	fmt.Printf("三参数指定类型与长度与容量%#v\n", names)
	fmt.Printf("获取容量与长度%#v，%#v\n", len(names), cap(names))
	//切片读取与修改
	names[0] = "update1"
	fmt.Printf("%#v\n", names)

	//添加元素 用append
	names = make([]string, 3)
	names[0] = "update0"
	names[1] = "update1"
	names[2] = "update2"
	names = append(names, "append1")
	fmt.Printf("%v\n", names)
	fmt.Printf("获取容量与长度%#v，%#v\n", len(names), cap(names))

	//遍历
	//长度函数  len
	//遍历1 for
	for i := 0; i < len(names); i++ {
		fmt.Printf("for len 遍历%v\n", names[i])
	}
	//遍历2 range
	for i, v := range names {
		fmt.Printf("for range 遍历%v,%v\n", i, v)
	}

	//copy 切片之间的值

	//长度相等
	//copy(目的，源)
	aSlice := []string{"a", "b"}
	bSlice := []string{"x", "y", "z"}
	copy(aSlice, bSlice)
	fmt.Printf("%v %v\n", aSlice, bSlice)

	//切片操作 =>数组，切片
	//方法1 nums[start:end]左包含右不包含  start end 可以不写，哪个不写就代表包含那一边所有，end 必须小于等于cap
	//方法2 nums[start:end:max] max可以不指定,不指定的话就到nums容量结尾
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	numChildren := nums[1:3]
	fmt.Printf("切片操作%v ,%v\n", nums, numChildren)

	//切片操作是共享内存空间，实际上nums numChildren在底层是同一个数组。
	//numChildren修改，同样会修改nums
	//numChildren的容量等于 nums的容量减去 start
	nums = make([]int, 5, 10)
	numChildren = nums[1:3]
	fmt.Printf("新切片长度与容量%v ,%v\n", len(numChildren), cap(numChildren))
	numChildren = append(numChildren, 100)
	numChildren[1] = 1000
	fmt.Printf("新切片修改%v ,%v\n", nums, numChildren)

	//数组切片操作
	//start <= end <=cap , new_cap=len-start
	numArray := [6]int{0, 1, 2, 3, 4, 5}
	arrayChildren := numArray[3:6]
	fmt.Printf("%#v ,%#v\n", numArray, arrayChildren)

	//nums[start:end:max]
	//start <= end <= max <= cap,new_cap=max-start
	numArray = [6]int{0, 1, 2, 3, 4, 5}
	arrayChildren = numArray[3:4:4]
	arrayChildren[0] = 2000 //没有生成新数组arrayChildren 更新导致numArray被更新
	fmt.Printf("%#v ,%#v\n", numArray, arrayChildren)
	arrayChildren = append(arrayChildren, 500) //生成新数组arrayChildren 更新不会导致numArray被更新
	fmt.Printf("%#v ,%#v\n", numArray, arrayChildren)
}
