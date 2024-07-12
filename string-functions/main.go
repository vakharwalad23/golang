package main

import (
	"fmt"
	s "strings"
)

var prt = fmt.Println

func main() {
	prt("Contains: ", s.Contains("test", "st"))
	prt("Count: ", s.Count("test", "s"))
	prt("HasPrefix: ", s.HasPrefix("test", "te"))
	prt("HasSuffix: ", s.HasSuffix("test", "st"))
	prt("Index: ", s.Index("test", "e"))
	prt("Join: ", s.Join([]string{"j", "o", "h", "n"}, "-"))
	prt("Repeat: ", s.Repeat("j", 5))
	prt("Replace: ", s.Replace("Booo", "o", "0", -1))
	prt("Replace: ", s.Replace("Booo", "o", "0", 1))
	prt("Split: ", s.Split("j-o-h-n", "-"))
	prt("ToLower: ", s.ToLower("JOHN"))
	prt("ToUpper: ", s.ToUpper("john"))
}
