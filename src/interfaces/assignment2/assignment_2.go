package main

import "fmt"

type Shape interface {
	getArea() float64
}

type Triangle struct {
	base   float64
	height float64
}

type Square struct {
	sideLength float64
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t Triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func PrintArea(s Shape) {
	fmt.Println(s.getArea())
}

func main() {
	s := Square{sideLength: 2}
	t := Triangle{base: 2, height: 2}
	PrintArea(t)
	PrintArea(s)
}
