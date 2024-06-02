package main

import (
	"errors"
	"fmt"
)

type TopError struct {
	msg string
	mid error
}

func (et TopError) Error() string {
	return et.msg
}

func (et TopError) Unwrap() error {
	return et.mid
}

type MiddleError struct {
	msg string
	bot error
}

func (em MiddleError) Error() string {
	return em.msg
}

func (em MiddleError) Unwrap() error {
	return em.bot
}

type BottomError struct {
	msg string
}

func (eb BottomError) Error() string {
	return eb.msg
}

func main() {
	bot := BottomError{msg: "bottom error"}
	mid := MiddleError{msg: "middle error", bot: bot}
	top := TopError{msg: "top error", mid: mid}

	if errors.Is(top, mid) {
		fmt.Println("Middle error is there!")
	}
	if !errors.Is(mid, top) {
		fmt.Println("Top error is not there!")
	}
	if errors.Is(top, bot) {
		fmt.Println("Bottom error is there!")
	}

	var err TopError
	if errors.As(top, &err) {
		fmt.Println("Top error is there!")
	}

	fmt.Printf("Error Chain: %w\n", top)

}
