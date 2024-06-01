package main

import "fmt"

func zeroval(n int) int {
	return n
}

func valptr(n *int) *int {
	return n
}

func main() {
	v1 := zeroval(5)
	fmt.Println(v1)

	val := 5
	v2 := valptr(&val)

	fmt.Println(v2)
	fmt.Println(&val)

	fmt.Println(*v2)
}
