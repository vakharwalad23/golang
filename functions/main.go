package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func multiReturn() (int, int) {
	return 1, 2
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	a, b := multiReturn()
	fmt.Println("multiReturn:", a, b)
	_, b = multiReturn()
	fmt.Println("multiReturn:", b)
}
