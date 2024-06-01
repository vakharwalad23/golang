package main

import "fmt"

func nums(nums ...int) {
	fmt.Print(nums, " -> ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	nums(1, 2)
	s := []int{1, 2, 3, 4}
	nums(s...)
}
