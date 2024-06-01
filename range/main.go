package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	for i, num := range nums {
		fmt.Println(i, num)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}
}
