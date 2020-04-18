package main

import "fmt"

func main()  {
	//值类型 a=b (在内存中申请内存新的空间，将A的值拷贝到b中)
	// int,float,point 数组类型，结构体类型--都是值类型
	//引用类型 a=b 是吧地址赋值给a
	//切片，映射，接口
	//验证方法 获取变量的的地址，与值

	//赋值《==》 形参，实参
	a:=0
	b:=make([]int,10)
	test1(a)
	fmt.Println(a)
	fmt.Printf("%v\n",&a)
	fmt.Printf("%v\n","########")
	test2(b)
	fmt.Println(b)
	fmt.Printf("%p\n",b)
}

func test1(n int)  {
	fmt.Printf("%v\n",&n)
	n=1
	
}

func test2(s []int)  {
	fmt.Printf("%p\n",s)
	s[0]=1
}

