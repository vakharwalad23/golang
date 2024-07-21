package main

import (
	"fmt"
	"time"
)

func main() {
	pln := fmt.Println

	now := time.Now()
	pln(now)

	time := time.Date(
		2004, 9, 23, 8, 5, 58, 651387237, time.UTC)
	pln(time)

	pln(time.Year())
	pln(time.Month())
	pln(time.Day())
	pln(time.Hour())
	pln(time.Minute())
	pln(time.Second())
	pln(time.Nanosecond())
	pln(time.Location())

	pln(time.Weekday())

	pln(time.Before(now))
	pln(time.After(now))
	pln(time.Equal(now))

	diff := now.Sub(time)
	pln(diff)

	pln(diff.Hours())
	pln(diff.Minutes())
	pln(diff.Seconds())
	pln(diff.Nanoseconds())

	pln(time.Add(diff))
	pln(time.Add(-diff))
}
