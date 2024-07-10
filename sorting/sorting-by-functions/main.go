package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"banana", "kiwi", "apple"}
	fmt.Println(fruits)

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		{name: "John", age: 30},
		{name: "Jack", age: 25},
		{name: "Rayn", age: 50},
	}

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(people)
}
