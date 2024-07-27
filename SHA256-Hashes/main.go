package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	str := "This is SHA256"

	h := sha256.New()
	hF := sha512.New()

	h.Write([]byte(str))
	hF.Write([]byte(str))

	bs := h.Sum(nil)
	bsF := hF.Sum(nil)

	fmt.Println(str)
	fmt.Printf("%x\n", bs)
	fmt.Printf("%x\n", bsF)
}
