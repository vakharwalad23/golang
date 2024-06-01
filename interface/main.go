package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
}

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (s square) area() float64 {
	return math.Pow(s.length, 2)
}

func (c circle) area() float64 {
	return (math.Pi * math.Pow(c.radius, 2))
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
}

func main() {
	c := circle{
		radius: 2.0,
	}
	s := square{
		length: 2.0,
	}

	measure(c)
	measure(s)
}
