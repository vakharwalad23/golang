package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"c", "a", "b"}
	fmt.Println("UnSorted: ", strs)
	slices.Sort(strs)
	fmt.Println("Sorted: ", strs)

	ints := []int{5, 7, 4}
	fmt.Println("Sorted: ", ints)
	slices.Sort(ints)
	fmt.Println("Sorted: ", ints)

	s := slices.IsSorted(ints)
	fmt.Println("IsSorted: ", s)
}
