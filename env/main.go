package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("PING", "1")
	fmt.Println("PING:", os.Getenv("PING"))
	fmt.Println("PONG:", os.Getenv("PONG"))

	fmt.Println()

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0], ":", pair[1])
	}
}
