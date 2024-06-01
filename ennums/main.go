package main

import "fmt"

type State int

const (
	Red = iota
	Green
	Blue
	Yellow
)

// By implementing the [fmt.Stringer](https://pkg.go.dev/fmt#Stringer)
// interface, values of `ServerState` can be printed out or converted
// to strings.
//
// This can get cumbersome if there are many possible values. In such
// cases the [stringer tool](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
// can be used in conjunction with `go:generate` to automate the
// process.
var stateName = map[State]string{
	Red:    "Red",
	Green:  "Green",
	Blue:   "Blue",
	Yellow: "Yellow",
}

func (s State) String() string {
	return stateName[s]
}

func transition(s State) State {
	switch s {
	case Red:
		return Red
	case Green:
		return Green
	case Blue:
		return Blue
	case Yellow:
		return Red
	default:
		return Red
	}
}

func main() {
	ns := transition(Blue)
	fmt.Println(ns)

	nst := transition(Yellow)
	ns2 := transition(nst)
	fmt.Println(ns2)
}
