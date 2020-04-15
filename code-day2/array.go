package main

import "fmt"

func main() {
	//数组定义1 【长度】类型，2【...】类型
	//数组的长度是固定，数组的长度是数组的一部分
	var names [3]string
	var signIns [3]bool
	var scores [3]float64
	fmt.Printf("%T\n", names)
	fmt.Printf("%T\n", signIns)
	fmt.Printf("%T\n", scores)

	fmt.Printf("%v\n", names)
	fmt.Printf("%v\n", signIns)
	fmt.Printf("%v\n", scores)

	//数组赋值方式
	names = [3]string{"a", "b", "c"}
	fmt.Printf("%#v\n", names)
	names = [3]string{1: "bbb"}
	fmt.Printf("%#v\n", names)
	names = [...]string{"aa", "bb", "cc"}
	fmt.Printf("%#v\n", names)

	//数组没有算术运算
	//数组的关系运算==  !=
	fmt.Println(names == [3]string{})
	fmt.Println(names == [3]string{"", "bbb", ""})

	//元素
	//访问￥ 修改 索引(从左到右 0，1，2，3，4...)
	fmt.Printf("%q\n", names[0])
	names[0] = "大雨"
	fmt.Printf("%q，%#v\n", names[0], names)

	//长度函数  len
	//遍历1 for
	for i := 0; i < len(names); i++ {
		fmt.Printf("for len 遍历%v\n", names[i])
	}
	//遍历2 range
	for _, v := range names {
		fmt.Printf("for range 遍历%v\n", v)
	}

	//定义多维数组
	d2 := [3][2]int{}
	fmt.Printf("%v\n", d2) //[2]int{0, 0}, [2]int{0, 0}, [2]int{0, 0}
	fmt.Printf("%v\n", d2[0])

}
