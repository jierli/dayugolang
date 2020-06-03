package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5
	areaIntf = sq1
	//fmt.Printf("areaintf的类型:%T,%v\n", areaIntf, areaIntf.side)
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("areaintf的类型:%T,%v\n", t, t.side)
	}
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("Circle的类型:%T,%v\n", u, u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}

	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("areaintf的类型:%T,%v\n", t, t)
	default:
		fmt.Printf("未知")
	}
}
