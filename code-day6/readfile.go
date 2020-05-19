package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file, _ := ioutil.ReadFile("os.go")
	fmt.Println(string(file))
}
