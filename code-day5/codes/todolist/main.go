package main

import (
	"fmt"
	"todolist/models"
)

/*
结构体可见范围
首字母必须大小，结构体名与 元素属性名
*/
func main() {
	var taskv2 models.Taskv2
	fmt.Println(taskv2)
	var anonyWrapperv2 models.AnonyWrapperv2
	fmt.Println(anonyWrapperv2)
	var anonyWrapperv3 models.AnonyWrapperv3
	fmt.Println(anonyWrapperv3)
}
