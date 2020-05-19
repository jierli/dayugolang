package main

import (
	"io"
	"os"
)

func main() {
	logFile1, _ := os.Create("./test/1.log")
	logFile2, _ := os.Create("./test/2.log")
	writer := io.MultiWriter(logFile1, logFile2, os.Stdout)
	writer.Write([]byte("hello world"))

	logFile1.Close()
	logFile2.Close()
}
