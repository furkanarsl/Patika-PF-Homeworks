package main

import (
	"fmt"
	"strings"
)

func main() {
	greetPrinter(createGreetInTurkish, "Furkan")
	greetPrinter(createGreetInEnglish, "John")
	greetPrinter(convertToUpperCase, "naber?")

	greetCreator := createGreetInTurkish
	greetPrinter(greetCreator, "Zeynep")

	//closures
	func(name string) {
		greeting := "Selam " + name
		fmt.Printf("%s\n", greeting)
	}("Yesim")

	closure := func(name string) {
		greeting := "Selam " + name
		fmt.Printf("%s\n", greeting)
	}
	closure("Fatma")
}

func createGreetInTurkish(name string) string {
	return "Selam " + name
}
func createGreetInEnglish(name string) string {
	return "Hi " + name
}

func convertToUpperCase(arg string) string {
	return strings.ToUpper(arg)
}

func greetPrinter(function func(it string) string, name string) {
	var greeting = function(name)
	fmt.Printf("%s\n", greeting)
}
