package main

import (
	"fmt"
	ctask "todolist/controllers/task"
	"todolist/models" //目录名
	"todolist/models/task"

	"github.com/imsilence/strutil"
	"github.com/jierli/dayupaktest"
	//引用包存到默认gopath
)

//导入目录名
//调用包名
//所以目录名与包名一致

func main() {
	fmt.Println(models.Name)
	fmt.Println(task.Name)
	fmt.Println(ctask.Name)
	task.Call()
	fmt.Println(task.GetVersion())
	fmt.Println(strutil.RandString())
	fmt.Println(dayupaktest.VarName)

}
