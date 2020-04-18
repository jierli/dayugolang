package main

import "fmt"

//递归
//n阶乘
//f(n)=n*f(n-1) 解法名--分治--大问题分解为多个相同的小问题--代码中实现：递归调用
//递归调用：函数直接或者间接调用自己，必须有结束条件

func f(n int64) int64 {
	if n == 0 { //结束条件  f(0)=1
		return 1
	}
	return n * f(n-1)
}

func main() {
	fmt.Println(f(4))
}
