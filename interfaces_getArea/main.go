package main

import "fmt"

type shape interface {
	getArea() float64
	getShapeType() string
}

type shapeSpecs struct {
	shapeName string
}

type square struct {
	sideLength float64
	shapeSpecs
}

type triangle struct {
	height float64
	base   float64
	shapeSpecs
}

func (t triangle) getArea() float64 {
	return (.5 * t.base) * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

// with composition we can extend square and triangle to print shapeName with only one 
// function
func (ss shapeSpecs) getShapeType() string {
	return ss.shapeName
}

func main() {
	s := square{sideLength: 2, shapeSpecs: shapeSpecs{shapeName: "Square"}}
	t := triangle{base: 2, height: 2, shapeSpecs: shapeSpecs{shapeName: "Triangle"}}

	printArea(s)
	printArea(t)

}

func printArea(s shape) {
	fmt.Println("Area of the", s.getShapeType(), ":", s.getArea())
}
