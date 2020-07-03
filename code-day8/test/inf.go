package main

import (
	"fmt"
	"io"
	"os"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

type Square1 struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}
func (sq *Square1) Area() float32 {
	return sq.side + sq.side
}
func main() {
	sq1 := new(Square)
	sq1.side = 5

	sq11 := new(Square1)
	sq11.side = 500

	var areaIntf Shaper
	areaIntf = sq1

	var areaIntf1 Shaper
	areaIntf1 = sq11

	fmt.Printf("The square has area: %f\n", areaIntf.Area())
	fmt.Printf("The square has area: %f\n", areaIntf1.Area())
	var w io.Writer
	w = os.Stdout
	fmt.Printf("%T,,,%#v", w, w)
	w.Write([]byte("hello"))
}
