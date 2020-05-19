package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	files, err := ioutil.ReadDir(".")
	fmt.Println(err)
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("D", file.Name(), file.ModTime())
		} else {
			fmt.Println("F", file.Name(), file.Size(), file.ModTime())
		}
	}
}
