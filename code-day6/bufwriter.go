package main

import (
	"bufio"
	"os"
)

func main() {
	file, _ := os.Open("./test/test.log")
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	writer.Write([]byte("abcdes\n"))
	writer.WriteString("1234567")

}
