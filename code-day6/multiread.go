package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	logFile1, _ := os.Open("./test/1.log")
	logFile2, _ := os.Open("./test/2.log")
	reader := io.MultiReader(logFile1, logFile2)
	buffer := make([]byte, 5)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(n, err, string(buffer[:n]))
	}

	logFile1.Close()
	logFile2.Close()
}
