package main

import (
	"fmt"
)

const (
	s string = "Hello"
)

func main() {
	fmt.Println(s)

	const d = 64e5

	fmt.Println(int64(d))
}
