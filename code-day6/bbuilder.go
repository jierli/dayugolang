package main

import (
	"bytes"
	"fmt"
)

func main() {
	buffer1 := bytes.NewBuffer([]byte("dayu,"))
	buffer2 := bytes.NewBufferString("123,")

	//写操作
	buffer1.Write([]byte("xyz"))
	buffer2.Write([]byte("789"))

	buffer1.WriteString("mn")
	buffer2.WriteString("55")

	//读操作
	ctx := make([]byte, 3)
	n, _ := buffer1.Read(ctx)
	fmt.Println(string(ctx[:n]))

	n, _ = buffer2.Read(ctx)
	fmt.Println(string(ctx[:n]))

	ctx, _ = buffer
}
