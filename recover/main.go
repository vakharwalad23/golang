package main

import "fmt"

func panicTrigger() {
	panic("Problem Occured!!")
}

func main() {

	defer func() {
		if r := recover(); r != nil {

			fmt.Println("Recovered from an Error:\n", r)
		}
	}()

	panicTrigger()

	fmt.Println(("After Panic"))
}
