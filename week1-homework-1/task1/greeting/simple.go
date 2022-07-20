package main

import "fmt"

func main() {
	// Basic
	fmt.Println("Selam ")
	greet("Furkan")

	// Functions with return
	var name = "Furkan"
	var greeting = createGreet(name)
	fmt.Printf("%s\n", greeting)

	// User input with Scanln
	var userName string
	fmt.Println("Please enter your name: ")
	fmt.Scanln(&userName)
	var greeting2 = createGreet(userName)
	fmt.Printf("%s\n", greeting2)

}

func greet(name string) {
	fmt.Printf("Selam %s\n", name)
}

func createGreet(name string) string {
	greeting := "Selam " + name
	return greeting
}
