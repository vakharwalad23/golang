package main

import (
	"fmt"
	"time"
)

func msg(m string) {
	for i := range 3 {
		fmt.Println(i, " : ", m)
	}
}

func main() {
	msg("Direct")

	go msg("Routine")

	go func(msg string) {
		fmt.Println("Anonymous function:", msg)
	}("Going")

	time.Sleep(time.Second)
	fmt.Println("Done")
}
