package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {

	//
	cmd := exec.Command("ping", "www.baidu.com")
	//bytes, err := cmd.Output()
	//fmt.Println(string(bytes), err)

	//同步管道输出
	output, err := cmd.StdoutPipe()
	cmd.Start()
	fmt.Println(err)
	io.Copy(os.Stdout, output)
	cmd.Wait()
}
