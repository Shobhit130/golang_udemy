package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

type shape interface {
	area()
}

func (s square) area() {
	fmt.Println(s.length * s.width)
}

func (c circle) area() {
	fmt.Println(math.Pi * (math.Pow(c.radius, 2)))
}

func info(s shape) {
	s.area()
}

func main() {
	s := square{
		length: 2,
		width:  3,
	}

	c := circle{
		radius: 3,
	}

	info(s)
	info(c)
}
