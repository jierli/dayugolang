package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, _ := os.Open("os.go")
	defer file.Close()
	ctx, _ := ioutil.ReadAll(file)
	fmt.Println(string(ctx))

}
