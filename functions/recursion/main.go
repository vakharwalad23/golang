package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(3))

	// Recursion with closures requires to be declared with a typed var explicitly before it’s defined
	var factorial func(n int) int

	factorial = func(n int) int {
		if n == 0 {
			return 1
		}

		return n * fact(n-1)
	}

	fmt.Println(factorial(4))
}
