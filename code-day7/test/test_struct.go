package main

import "fmt"

type Author struct {
	id   int
	name string
	addr string
	tel  string
	desc string
}

var kk Author
var pkk *Author

var dayus Author = Author{}
var dayus2 Author = Author{
	1001,
	"dayu",
	"beijing",
	"186183",
	"dba",
}

var authornew *Author = new(Author)

func main() {
	fmt.Printf("%T   :%#v:    %p\n", kk, kk, kk)
	fmt.Printf("%T   :%#v:    %p\n", pkk, pkk, pkk)

	fmt.Printf("%T   :%#v:    %p\n", dayus, dayus, dayus)
	fmt.Printf("%T   :%#v:    %p\n", dayus2, dayus2, dayus2)

	fmt.Printf("%T   :%#v:    %p\n", authornew, authornew, authornew)

	dayuNew := NewAuthor(1111, "dayu2", "xiamen", "177", "yunwei")
	fmt.Printf("%T   :%#v:    %p\n", dayuNew, dayuNew, dayuNew)
	dayuNew.desc = "zongjian"
	fmt.Printf("%T   :%#v:    %p\n", dayuNew, dayuNew, dayuNew)
	authornew.desc = "haha"
	fmt.Printf("%T   :%#v:    %p\n", authornew, authornew, authornew)

}

func NewAuthor(id int, name, addr, tel, desc string) *Author {
	return &Author{id, name, addr, tel, desc}
}
