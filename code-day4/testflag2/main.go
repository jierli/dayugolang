package main

import (
	"flag"
	"fmt"
)

/*
var (
	host string
	port int
	h    bool
	help bool
)
*/

//命令行中的参数进行解析-》指定的命令参数-》赋值到变量 用flag包
func main() {
	//-H host -P port

	//String（命令行指定输入参数名，默认值，帮助说明）
	host := flag.String("H", "127.0.0.1", "连接地址")
	port := flag.Int("P", 22, "连接端口")
	h := flag.Bool("h", false, "连接端口")
	help := flag.Bool("help", false, "连接端口")

	flag.Usage = func() {
		fmt.Println("usage: -h")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *h || *help {
		flag.Usage()
		return
	}

	fmt.Printf("%#V,%#V,%#V,%#V\n", host, port, h, help)
	fmt.Println(*host, *port, *h, *help)

	fmt.Println(flag.NArg)
	fmt.Printf("%#v", flag.Args())
}
