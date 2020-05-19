package main

import (
	"fmt"
	"os"
)

func main() {
	logFile1, _ := os.Open("./test/1.log")
	defer logFile1.Close()

	buffer := make([]byte, 10)

	ctx := make([]byte, 1024*1024)
	for {
		n, err := logFile1.Read(buffer)
		if err != nil {
			break
		}
		ctx = append(ctx, buffer[:n]...)
	}
	fmt.Println(string(ctx))

}
