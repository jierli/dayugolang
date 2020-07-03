package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main()  {

	sttr:= strings.NewReader("cat|1|filename|")
	fmt.Printf("read cnt err: %#v\n", sttr)
	//reader := bufio.NewReader(os.Stdin)
	reader := bufio.NewReader(sttr)
	//fmt.Printf("read cnt err: %#v\n", reader)
	cntText, err := reader.ReadString('|')
	fmt.Printf("read cnt err: %#v,%s\n", cntText,err)
	cnt := cntText[:len(cntText)-1]
	fmt.Printf("read cnt err: %#v,%#v\n", cnt)

}
