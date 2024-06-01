package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprint("Base: ", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			num: 23,
		},
		str: "Hello",
	}
	fmt.Println(co.num, co.str)
	fmt.Println(co.base.num)
	fmt.Println(co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println(d.describe())
}
