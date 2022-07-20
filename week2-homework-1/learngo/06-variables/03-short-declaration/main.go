package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE 1: Short Declare
//
//  Declare and then print four variables using
//  the short declaration statement.
//
// EXPECTED OUTPUT
//  i: 314 f: 3.14 s: Hello b: true
// ---------------------------------------------------------
// EXERCISE 2: Multiple Short Declare
//
//  Declare two variables using multiple short declaration
//
// EXPECTED OUTPUT
//  14 true
// ---------------------------------------------------------
// EXERCISE 3: Multiple Short Declare #2
//
//  1. Declare two variables using multiple short declaration
//  2. `a` variable's value should be 42
//  3. `c` variable's value should be "good"
//
// EXPECTED OUTPUT
//  42 good
// ---------------------------------------------------------
// EXERCISE 4: Short With Expression
//
// 	1. Short declare a variable named `sum`
//  2. Initialize it with an expression by adding 27 and 3.5
//
// EXPECTED OUTPUT
//  30.5
// ---------------------------------------------------------
// EXERCISE 5: Short Discard
//
// 	1. Short declare two bool variables
//     (use multiple short declaration syntax)
//  2. Initialize both variables to true
//  3. Change your declaration and
//     discard the 2nd variable's value
//     using the blank-identifier
//  4. Print only the 1st variable
//
// EXPECTED OUTPUT
//  true
// ---------------------------------------------------------
// EXERCISE 6: Redeclare
//
// 	1. Short declare two int variables: age and yourAge
//     (use multiple short declaration syntax)
//  2. Short declare a new float variable: ratio
//     And, change the 'age' variable to 42
//     (! You should use redeclaration)
//  4. Print all the variables
//
// EXPECTED OUTPUT
//  42, 20, 3.14

func main() {
	//Exercise 1
	// ADD YOUR DECLARATIONS HERE
	i := 314
	f := 3.14
	s := "Hello"
	b := true
	fmt.Println(
		"i:", i,
		"f:", f,
		"s:", s,
		"b:", b,
	)
	// Exercise 2
	a, b := 14, true
	fmt.Println(a, b)

	//Exercise 3
	c, d := 42, "good"
	fmt.Println(c, d)

	// Exercise 4
	sum := 27 + 3.5
	fmt.Println(sum)

	//Exercise 5
	b1, _ := true, true
	fmt.Println(b1)

	//Exercise 6
	age, yourAge := 1, 20
	age, ratio := 42, 3.14

	fmt.Println(age, yourAge, ratio)
}
