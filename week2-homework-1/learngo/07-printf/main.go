package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE 1: Print Your Age
//  Print your age using Printf
// EXPECTED OUTPUT
//  I'm 30 years old.
// NOTE
//  You should change 30 to your age, of course.
// ---------------------------------------------------------
// EXERCISE 2: Print Your Name and LastName
//  Print your name and lastname using Printf
// EXPECTED OUTPUT
//  My name is Inanc and my lastname is Gumus.
// BONUS
//  Store the formatting specifier (first argument of Printf)
//    in a variable.
//  Then pass it to printf
// ---------------------------------------------------------
// EXERCISE 3: False Claims
//  Use printf to print the expected output using a variable.
// EXPECTED OUTPUT
//  These are false claims.
// ---------------------------------------------------------
// EXERCISE 4: Print the Temperature
//  Print the current temperature in your area using Printf
// NOTE
//  Do not use %v verb
//  Output "shouldn't" be like 29.500000
// EXPECTED OUTPUT
//  Temperature is 29.5 degrees.
// ---------------------------------------------------------
// EXERCISE 5: Double Quotes
//  Print "hello world" with double-quotes using Printf
//  (As you see here)
// NOTE
//  Output "shouldn't" be just: hello world
// EXPECTED OUTPUT
//  "hello world"
// ---------------------------------------------------------
// EXERCISE 6: Print the Type
//  Print the type and value of 3 using fmt.Printf
// EXPECTED OUTPUT
//  Type of 3 is int
// Exercise 7-9: Same as 6 but with different types.
// ---------------------------------------------------------
// EXERCISE 10: Print Your Fullname
//  1. Get your name and lastname from the command-line
//  2. Print them using Printf
// EXAMPLE INPUT
//  Inanc Gumus
// EXPECTED OUTPUT
//  Your name is Inanc and your lastname is Gumus.
// ---------------------------------------------------------
func main() {
	// Since Exercise#10 requires command line arguments program will error if they are not provided

	//Exercise 1
	fmt.Printf("I'm %v years old.\n", 23)

	//Exercise 2
	hello := "My name is %s and my lastname is %s.\n"
	fmt.Printf(hello, "Furkan", "Arslan")

	//Exercise 3
	// UNCOMMENT THE FOLLOWING CODE
	// AND DO NOT CHANGE IT AFTERWARDS
	tf := false
	// TYPE YOUR CODE HERE
	fmt.Printf("There are %t claims.\n", tf)

	//Exercise 4
	fmt.Printf("Temperature is %.1f degrees.\n", 29.5)

	//Exercise 5
	fmt.Printf("%q\n", "hello world")

	//Exercise 6
	fmt.Printf("Type of %d is %[1]T\n", 3)

	//Exercise 7
	fmt.Printf("Type of %f is %[1]T\n", 3.14)

	//Exercise 8
	fmt.Printf("Type of %s is %[1]T\n", "hello")

	//Exercise 9
	fmt.Printf("Type of %t is %[1]T\n", true)

	// Exercise 10
	name, lastName := os.Args[1], os.Args[2]

	c := "Your name is %s and your lastname is %s.\n"
	fmt.Printf(c, name, lastName)

}
