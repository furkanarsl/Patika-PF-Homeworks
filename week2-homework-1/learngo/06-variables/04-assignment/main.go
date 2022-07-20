package main

import (
	"fmt"
	"path"
)

// ---------------------------------------------------------
// EXERCISE 1: Make It Blue
//  1. Change `color` variable's value to "blue"
//  2. Print it
//
// EXPECTED OUTPUT
//  blue
// ---------------------------------------------------------
// EXERCISE 2: Variables To Variables
//
//  1. Change the value of `color` variable to "dark green"
//  2. Do not assign "dark green" to `color` directly
//     Instead, use the `color` variable again
//     while assigning to `color`
//  3. Print it
// RESTRICTIONS
//  WRONG ANSWER, DO NOT DO THIS:
//  `color = "dark green"`
// HINT
//  + operator can concatenate string values
//  Example:
//  "h" + "e" + "y" returns "hey"
// EXPECTED OUTPUT
//  dark green
// ---------------------------------------------------------
// EXERCISE 3: Assign With Expressions
//  1. Multiply 3.14 with 2 and assign it to `n` variable
//  2. Print the `n` variable
// HINT
//  Example: 3 * 2 = 6
//  * is the product operator (it multiplies numbers)
// EXPECTED OUTPUT
//  6.28
// ---------------------------------------------------------
// EXERCISE 4: Find the Rectangle's Perimeter
//  1. Find the perimeter of a rectangle
//     Its width is  5
//     Its height is 6
//  2. Assign the result to the `perimeter` variable
//  3. Print the `perimeter` variable
// HINT
//  Formula = 2 times the width and height
// EXPECTED OUTPUT
//  22
// ---------------------------------------------------------
// EXERCISE 5: Multi Assign
//  1. Assign "go" to `lang` variable
//     and assign 2 to `version` variable
//     using a multiple assignment statement
//  2. Print the variables
//
// EXPECTED OUTPUT
//  go version 2
// ---------------------------------------------------------
// EXERCISE 6: Multi Assign #2
//  1. Assign the correct values to the variables
//     to match to the EXPECTED OUTPUT below
//  2. Print the variables
// HINT
//  Use multiple Println calls to print each sentence.
//
// EXPECTED OUTPUT
//  Air is good on Mars
//  It's true
//  It is 19.5 degrees
// ---------------------------------------------------------
// EXERCISE 7: Multi Short Func
// 	1. Declare two variables using multiple short declaration syntax
//  2. Initialize the variables using `multi` function below
//  3. Discard the 1st variable's value in the declaration
//  4. Print only the 2nd variable
// NOTE
//  You should use `multi` function
//  while initializing the variables
// EXPECTED OUTPUT
//  4
// ---------------------------------------------------------
// EXERCISE 8: Swapper
//  1. Change `color` to "orange"
//     and `color2` to "green" at the same time
//     (use multiple-assignment)
//  2. Print the variables
// EXPECTED OUTPUT
//  orange green
// ---------------------------------------------------------
// EXERCISE 9: Swapper #2
//  1. Swap the values of `red` and `blue` variables
//  2. Print them
// EXPECTED OUTPUT
//  blue red
// ---------------------------------------------------------
// EXERCISE: Discard The File
//  1. Print only the directory using `path.Split`
//  2. Discard the file part
// RESTRICTION
//  Use short declaration
// EXPECTED OUTPUT
//  secret/
// ---------------------------------------------------------

func main() {
	//Exercise 1
	color := "green"
	color = "blue"
	fmt.Println(color)

	// Exercise 2
	color2 := "green"
	color2 = "dark " + color2
	fmt.Println(color2)

	//Exercise 3
	n := 0.
	n = 3.14 * 2
	fmt.Println(n)

	//Exercise 4
	var (
		perimeter     int
		width, height = 5, 6
	)
	perimeter = 2 * (width * height)
	fmt.Println(perimeter)

	//Exercise 5
	var (
		lang    string
		version int
	)
	lang, version = "go", 2
	fmt.Println(lang, "version", version)

	//Exercise 6
	var (
		planet string
		isTrue bool
		temp   float64
	)
	planet, isTrue, temp = "Mars", true, 19.5
	fmt.Println("Air is good on", planet)
	fmt.Println("It's", isTrue)
	fmt.Println("It is", temp, "degrees")

	//Exercise 7
	_, v2 := multi()
	fmt.Println(v2)

	//Exercise 8
	c1, c2 := "red", "blue"
	c1, c2 = "orange", "green"
	fmt.Println(c1, c2)

	//Exercise 9
	red, blue := "red", "blue"
	red, blue = blue, red
	fmt.Println(red, blue)

	// Exercise 10

	dir, _ := path.Split("secret/file.txt")
	fmt.Println(dir)
}

func multi() (int, int) {
	return 5, 4
}
