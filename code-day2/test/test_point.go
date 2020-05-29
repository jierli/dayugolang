package main

import "fmt"

var (
	pointer01 *int
	pointer02 *float64
	pointer03 *string
)

var (
	age   int     = 30
	heigh float64 = 1.7
	motto string  = "少年"
)

func main() {
	fmt.Printf("%T  %T  %T\n", pointer01, pointer02, pointer03)
	fmt.Printf("%t  %t  %t", pointer01 == nil, pointer02, pointer03)
	pointer01, pointer02, pointer03 = &age, &heigh, &motto
	pointer04, pointer05, pointer06 := &age, &heigh, &motto
	pointer07, pointer08, pointer09 := new(int), new(float64), new(string)

	fmt.Println(pointer01, pointer02, pointer03)
	fmt.Println(*pointer01, *pointer02, *pointer03)
	fmt.Println(*pointer04, *pointer05, *pointer06)
	*pointer04 = 33
	fmt.Println(*pointer01, *pointer02, *pointer03)
	fmt.Println(*pointer04, *pointer05, *pointer06)
	fmt.Printf("%v %v %q\n", pointer07, pointer08, pointer09)
	fmt.Printf("%#v %#v %q\n", *pointer07, *pointer08, *pointer09)

	age += 1
	heigh = 1.77
	motto = "我是大雨"
	fmt.Println(*pointer01, *pointer02, *pointer03)
	fmt.Println(*pointer04, *pointer05, *pointer06)
}
