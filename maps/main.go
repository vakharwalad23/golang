package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]string)

	m["foo"] = "bar"
	m["ping"] = "pong"

	fmt.Println(m)

	v1 := m["foo"]
	fmt.Println(v1)

	// Non existant val
	v3 := m["non"]
	fmt.Println(v3)

	fmt.Println(len(m))

	delete(m, "ping")
	fmt.Println(m)

	clear(m)
	fmt.Println(m)

	v, prs := m["foo"]
	fmt.Println(v)
	fmt.Println(prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)

	o := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, o) {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}
}
