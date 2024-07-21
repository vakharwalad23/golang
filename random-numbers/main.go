package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	pln := fmt.Println
	pl := fmt.Print
	pl(rand.IntN(100), ", ")
	pl(rand.IntN(100), "\n")

	pln(rand.Float64())

	pl((rand.Float64()*5)+5, ", ")
	pl((rand.Float64()*5)+5, "\n")

	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	pl(r2.IntN(100), ", ")
	pl(r2.IntN(100), "\n")

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	pl(r3.IntN(100), ", ")
	pl(r3.IntN(100), "\n")
}
