package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println(net.JoinHostPort("120.0.0.1", "6666"))
	net.JoinHostPort("120.0.0.1", "6666")
	log.Panicf(net.JoinHostPort("120.0.0.1", "6666"))
	test := bytes.NewBufferString("dayu|aaa|vvv|")
	bufio.NewReader(test)
	log.Fatal()

}
