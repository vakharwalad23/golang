package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("Hello\nFrom\nGo")
	err := os.WriteFile("dump.txt", d1, 0644)
	check(err)

	// Open a file for writing to get more granular control
	f, err := os.Create("dump1.txt")
	check(err)

	// It is idiomatic to close the file immediately after opening
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n, err := f.Write(d2)
	check(err)

	fmt.Printf("Bytes written %d\n",n)

	n1, err := f.WriteString("Hello \nGo World")
	check(err)
	fmt.Printf("Bytes written %d\n",n1)

	f.Sync()

	// Buffered writer from buffio
	w := bufio.NewWriter(f)
	n2, err := w.WriteString("\nbuffered\n")
	check(err)
	fmt.Printf("Bytes written %d\n",n2)

	w.Flush()
}