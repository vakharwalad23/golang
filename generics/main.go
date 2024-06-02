package main

import "fmt"

type Number interface {
	int | float64
}

func sumInt(m map[string]int) int {
	var s int
	for _, val := range m {
		s += val
	}
	return s
}

func sumFloat(m map[string]float64) float64 {
	var s float64
	for _, val := range m {
		s += val
	}
	return s
}

// With Generics

func sumIntFloat[K comparable, V int | float64](m map[K]V) V {
	var s V
	for _, val := range m {
		s += val
	}
	return s
}

func sumInFl[K comparable, V Number](m map[K]V) V {
	var s V
	for _, val := range m {
		s += val
	}
	return s
}

func main() {
	ints := map[string]int{"a": 1, "b": 2, "c": 3}
	floats := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}

	fmt.Println(sumInt(ints))
	fmt.Println(sumFloat(floats))

	fmt.Println(sumIntFloat[string, int](ints))
	fmt.Println(sumIntFloat[string, float64](floats))

	fmt.Println(sumIntFloat(ints))
	fmt.Println(sumIntFloat(floats))

	fmt.Println(sumInFl(ints))
	fmt.Println(sumInFl(floats))
}
