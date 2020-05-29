package main

import "fmt"

type Point struct {
	X, Y float64
}

func (p Point) Show() {
	fmt.Printf("Position{%v, %v}\n", p.X, p.Y)
}

func (p *Point) MoveTo(x float64, y float64) {
	p.X, p.Y = x, y
}

func main() {
	p := Point{4.5, 6.7}

	myShow := Point.Show
	// myMoveTo := Point.MoveTo // Not OK! Point.MoveTo undefined (type Point has no method MoveTo)
	myMoveTo := (*Point).MoveTo
	p.Show()
	fmt.Printf("Type of myShow: %T\n", myShow)
	fmt.Printf("Type of myMoveTo: %T\n", myMoveTo)

	myShow(p)
	myMoveTo(&p, 9.9, 10.1)
	myShow(p)
}
