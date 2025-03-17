package Model

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float32
}

type Rectangle struct {
	Width, Height float32
}

func (r Rectangle) area() float32 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float32
}

func (c Circle) area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func ShowRectangleArea(r Rectangle) {
	area := r.Width * r.Height
	fmt.Println(area)
}

func ShowCircleArea(c Circle) {
	area := math.Pi * c.Radius * c.Radius
	fmt.Println(area)
}

func ShowGeometry(g geometry) {
	fmt.Println(g.area())
}
