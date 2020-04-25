package task

import "fmt"

var Name = "models task test1111"

//init函数在包导入初始化数据并且执行
func init() {
	fmt.Println("init in task.go")
}

func Call() {
	fmt.Println("call models task func")
}

//属性
//首字母大写包外可见，包外可修改
//首字母小写私有属性包外不能调用，不能修改，如果需要对外提供只读变更时，通过公有函数return 私有变量。
//例如
var version = "version 1.1.1"

func GetVersion() string {
	return version
}
