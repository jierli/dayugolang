package main

import (
	"log"
	"os"
)

func main()  {
	//设置格式	//SetFlags
	log.SetFlags(log.Flags()|log.Ldate|log.Lshortfile)
	//设置前缀 //SetPrefix
	log.SetPrefix("main: ")

	log.Println("log 11111")

	//不可恢复性错误日志可以用Fatal （os.Exit(1)） 退出
	//log.Fatalln("Fatal log")

	//可恢复性错误可以用Panic打印日志 ，引用放可以用reconver 恢复执行
	//log.Panicln("panic 日志")

	//标准输出 自定义日志输出格式log.New
	logger:=log.New(os.Stdout,"logger:",log.Flags())
	logger2:=log.New(os.Stdout,"logger2:",log.Flags())
	logger.Println("logger 111111")
	logger2.Println("logger2 11111")

	log.Println("log 22222222")
	//标准输入 fmt.Scan 标准输出 fmt.Println
}