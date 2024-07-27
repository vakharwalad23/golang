package main

import (
	b64E "encoding/base64"
	"fmt"
)

func main() {
	data := "abcde!@#$123()~@"

	sEnc := b64E.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := b64E.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))

	uEnc := b64E.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(string(uEnc))
	uDec, _ := b64E.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
