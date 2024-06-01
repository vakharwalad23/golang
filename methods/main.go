package main

import (
	"fmt"
	"math"
)

type Square struct {
	Length float64
}

func (s Square) area() float64 {
	return math.Pow(s.Length, 2)
}

func (s *Square) perim() float64 {
	return 4 * s.Length
}

func main() {
	s := Square{
		Length: 2.5,
	}
	fmt.Println(s.area())
	fmt.Println(s.perim())

	s1 := &s
	fmt.Println(s1.area())
	fmt.Println(s1.perim())
}
