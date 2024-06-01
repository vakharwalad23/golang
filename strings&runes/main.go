// A Go string is a read-only slice of bytes. The language
// and the standard library treat strings specially - as
// containers of text encoded in [UTF-8]
// In other languages, strings are made of "characters".
// In Go, the concept of a character is called a `rune` - it's
// an integer that represents a Unicode code point.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello"

	// Since strings are equivalent to []byte, this will produce the length of the raw bytes stored within
	fmt.Println("Length", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for i, runeval := range s {
		fmt.Printf("%#U starts at %d\n", runeval, i)
	}

	fmt.Println("Decode using DecodeRuneInString from unicode/utf8 package.")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
	}
}
