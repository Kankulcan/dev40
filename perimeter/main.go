package main

import "math"

type Shape interface {
	area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct{
	Radius float64
}
type Triangle struct{
	height float64
	base float64
}


func main() {

}

func perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (t Triangle) area() float64 {
	return (t.base*t.height) *0.5
}