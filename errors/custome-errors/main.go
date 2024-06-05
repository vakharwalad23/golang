package main

import (
	"errors"
	"fmt"
)

type cutomError struct {
	arg     int
	message string
}

func (e *cutomError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func foo(arg int) (int, error) {
	if arg == 23 {
		return -1, &cutomError{arg, "Yo it is 23"}
	}
	return arg, nil
}

func main() {
	_, error := foo(23)
	var cutomError *cutomError
	if errors.As(error, &cutomError) {
		fmt.Println("Arg:", cutomError.arg, " Message:", cutomError.message)
	} else {
		fmt.Println("No match")
	}
}
