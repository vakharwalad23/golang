package main

import "fmt"

func closure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := closure()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newNextInt := closure()
	fmt.Println(newNextInt())
	fmt.Println(newNextInt())
}
