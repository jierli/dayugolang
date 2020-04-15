package main

import "fmt"

func main() {
	//map定义map[keyType]valueType
	//定义成绩map
	var scores map[string]int
	fmt.Printf("%T,    %#v\n", scores, scores)

	//初始化
	//map[string]int
	//字面量
	scores = map[string]int{}
	fmt.Printf("%#v\n", scores)
	scores = map[string]int{"a": 100, "b": 90, "c": 80}
	fmt.Printf("%#v\n", scores)

	//make
	scores = make(map[string]int)
	fmt.Printf("%T,%#v\n", scores, scores)

	//操作
	// 增删改查
	scores = map[string]int{"a": 100, "b": 90, "c": 80}
	//查
	fmt.Printf("key:a=%#v\n", scores["a"])
	//查2  判断key是否存在,使用两个返回值
	v, ok := scores["xx"]
	fmt.Println(v, ok)

	//改 增
	scores["xx"] = 99
	scores["a"] = 89
	fmt.Println(scores)

	//删除
	delete(scores, "xx")
	fmt.Println(scores)

	//遍历
	for k, i := range scores {
		fmt.Printf("%v,%v\n", k, i)
	}
}
