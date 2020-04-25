package main

import (
	"fmt"
	"time"
)

//声明存储todolist
var todolist = make(map[string][]string)

func createTask() {
	var taskID, taskName, taskExecTime string
	fmt.Println("请输入任务ID：")
	fmt.Scan(&taskID)
	fmt.Println("请输入任务Name：")
	fmt.Scan(&taskName)
	fmt.Println("请输入任务计划执行时间：")
	fmt.Scan(&taskExecTime)
	CreteTime := time.Now().Format("2006-01-02 15:04")
	taskState := "未开始"
	taskMap[taskID] = []string{
		taskName, taskExecTime, CreteTime, taskState,
	}
}

func main() {
	createTask()
}
