package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 1 {
		return 0, errors.New("Error!")
	}
	return arg, nil
}

var ErrOutOfPage = fmt.Errorf("Out of page")
var ErrOutOfPower = fmt.Errorf("Out of power")

func printPage(page int) error {
	if page == 0 {
		return ErrOutOfPage
	} else if page == 1 {
		return ErrOutOfPower
	}
	return nil
}

func main() {
	for i := range 3 {
		if val, err := f(i); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(val)
		}
	}

	for i := range 3 {
		if err := printPage(i); err != nil {
			if errors.Is(err, ErrOutOfPage) {
				fmt.Println("Add more page")
			} else if errors.Is(err, ErrOutOfPower) {
				fmt.Println("Connect to power")
			} else {
				fmt.Println("Something went wrong!")
			}
			continue
		}
		fmt.Println("Printed page", i)
	}
}
