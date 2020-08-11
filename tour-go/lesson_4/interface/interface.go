package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}
type rectangle struct {
	height,width float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.height*r.width
}

func (c circle) area() float64 {
	return math.Pi*c.radius*c.radius
}

func (r rectangle) perimeter()  float64  {
	return 2*r.height + 2*r.width
}

func (c circle) perimeter() float64 {
	return math.Pi*2*c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rectangle{3,4}
	c := circle{4}

	measure(r)
	measure(c)
}