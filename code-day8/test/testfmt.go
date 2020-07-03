package main

import (
	"fmt"
	"log"
	"os"
)

func main()  {
	a:=fmt.Sprint("aaaaaaa")
	b:="testbbbb"
	fmt.Printf("%#v,%T......%#v,%T",a,a,b,b)

	logfile, _ := os.OpenFile("testfmt.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	defer logfile.Close()

	// Todo: 带缓冲的流
	log.SetOutput(logfile)
	log.Printf(a)
	log.Printf(a)
}