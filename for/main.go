package main

import "fmt"

func main() {
	i := 1
	for i <= 2 {
		fmt.Println(i)
		i++
	}

	for j := 0; j <= 2; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("Range", i)
	}

	for {
		fmt.Println("Hello")
		break
	}

	for i := range 6 {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
