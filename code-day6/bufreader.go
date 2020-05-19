package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("./test/user.txt")
	defer file.Close()

	reader := bufio.NewReader(file)

	ctx, _ := reader.ReadBytes(',')
	fmt.Println(string(ctx))

	ctx, _ = reader.ReadBytes(',')
	fmt.Println(string(ctx))

	ctx, prefix, _ := reader.ReadLine()
	fmt.Println(string(ctx), prefix)
}
