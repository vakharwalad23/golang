package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func newPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func main() {
	fmt.Println(Person{Name: "John", Age: 19})
	fmt.Println(Person{"Doe", 19})
	fmt.Println(Person{Name: "Bob"})
	fmt.Println(&Person{Name: "Jack", Age: 19})

	person := newPerson("Ping", 19)
	fmt.Println(person)

	animal := struct {
		Name string
		wild bool
	}{
		Name: "Dog",
		wild: false,
	}
	fmt.Println(animal)
}
