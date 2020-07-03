package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	ctx := make([]byte, 1024)
	test := bytes.NewBufferString("dayu|aaa|vvv|")
	test1 := bytes.NewReader(ctx)
	fmt.Println(test1)
	fmt.Println(test.ReadString('|'))
	fmt.Println(test.ReadString('|'))
	fmt.Println(test.ReadString('|'))
	ctx = make([]byte, 1024)
	fmt.Println(string(ctx))
	n, err := test.Read(ctx)
	fmt.Println(n, err)
	fmt.Println(string(ctx))
	fmt.Println(strconv.FormatFloat(1.335555, 'f', 5, 64))
	a := struct {
		a string
		b int
	}{"a", 1}
	fmt.Printf("%v,%T", a, a)

	createTask := func(i int) func() interface{} {
		return func() interface{} {
			return i
		}
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("%T,%c\n", createTask(i))
	}



	type TimeHandler struct {
	}
	func (h *TimeHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
		now := time.Now().Format("2006-01-02 15:04:05")
		io.WriteString(response, now)
	}
	fmt.Printf("%T,%#v", &TimeHandler{})
}
