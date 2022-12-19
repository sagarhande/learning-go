package main

import (
	"fmt"
	"reflect"
)

type shape interface {
	getArea() float64
}

type squere struct {
	side float64
}
type triangle struct {
	base   float64
	height float64
}

// func main() {

// }

// Interface function
func printArea(s shape) {
	t := reflect.TypeOf(s)
	fmt.Println("Area of a "+t.Name()+" is :", s.getArea())
}

func (s squere) getArea() float64 {
	return s.side * s.side
}

func (t triangle) getArea() float64 {

	return 0.5 * t.base * t.height
}
