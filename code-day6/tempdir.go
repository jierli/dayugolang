package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	//创建临时文件夹
	//fmt.Println(ioutil.TempDir("./test", "log"))
	dir, _ := ioutil.TempDir("./test", "log")
	file, _ := os.Create(filepath.Join(dir, "4.log"))
}
