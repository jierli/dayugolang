package main

import (
	"flag"
	"fmt"
	"strings"
)

var sep = flag.String("s", "##", "name")

var num = flag.Int("n", 0, "age")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	fmt.Println("age=", *num)
}

//例子go run flagpage.go  -s % abbc nn ccc -n 1100
