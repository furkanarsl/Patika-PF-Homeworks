package main

import "fmt"

func main() {
	person := Person{"Furkan"}
	var greeting = person.greet()
	fmt.Printf("%v\n", greeting)
}

type Person struct {
	name string
}

func (p Person) greet() string {
	return "Selam " + p.name
}
