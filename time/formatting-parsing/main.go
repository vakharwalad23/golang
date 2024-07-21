package main

import (
	"fmt"
	"time"
)

func main() {
	pln := fmt.Println

	t := time.Now()
	pln(t.Format(time.RFC3339))

	t1, _ := time.Parse(time.RFC3339, "2024-07-21T18:24:23+05:30")
	pln(t1)

	pln(t.Format("3:04PM"))
	pln(t.Format("Mon Jan _2 15:04:05 2006"))
	pln(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	pln(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	pln(e)
}
