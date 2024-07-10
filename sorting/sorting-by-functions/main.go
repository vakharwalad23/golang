package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Person struct {
	name string
	age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.name, p.age)
}

func main() {
	fruits := []string{"banana", "kiwi", "apple"}
	fmt.Println(fruits)

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	people := []*Person{
		{name: "John", age: 30},
		{name: "Jack", age: 25},
		{name: "Rayn", age: 50},
	}

	slices.SortFunc(people, func(a, b *Person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(people)
}
