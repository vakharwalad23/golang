package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Entering to the new world!")

	os.Exit(1)
}
