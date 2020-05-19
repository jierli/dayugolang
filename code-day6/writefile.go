package main

import (
	"io/ioutil"
	"os"
)

func main() {
	ioutil.WriteFile("./test/user.txt", []byte("i am dayu"), os.ModePerm)
}
