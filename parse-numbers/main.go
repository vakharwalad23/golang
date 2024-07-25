package main

import (
	"fmt"
	"strconv"
)

func main() {
	pln := fmt.Println

	f, _ := strconv.ParseFloat("2.456", 64)
	pln(f)

	i, _ := strconv.ParseInt("23", 0, 64)
	pln(i)

	h, _ := strconv.ParseInt("0x1c8", 0, 64)
	pln(h)

	u, _ := strconv.ParseUint("234", 0, 64)
	pln(u)

	a, _ := strconv.Atoi("234")
	pln(a)

	_, e := strconv.Atoi("wait")
	pln(e)

}
