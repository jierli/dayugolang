package main

import "fmt"

//空接口
type EmptyIntf interface {
}

//空接口实现不同类型数据输入
func PrintType(v interface{}) {
	switch value := v.(type) {
	case int:
		fmt.Println("int", value)
	case []int:
		fmt.Println("[]int", value)
	case [3]int:
		fmt.Println("[3]int", value)
	case map[string]string:
		fmt.Println("map", value)
	default:
		fmt.Println("未知")
	}
}

func main() {
	var emptyIntf EmptyIntf
	emptyIntf = 1
	emptyIntf = "test"
	fmt.Println(emptyIntf)
	//匿名空接口
	var emptyIntf2 interface{}
	emptyIntf2 = 1
	emptyIntf2 = "test"

	fmt.Println(emptyIntf2)

	PrintType(1)
	PrintType("test")

}
