package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE 1: Convert and Fix
//  Fix the code by using the conversion expression.
// EXPECTED OUTPUT
//  15.5
// ---------------------------------------------------------
// EXERCISE 2: Convert and Fix #2
//  Fix the code by using the conversion expression.
// EXPECTED OUTPUT
//  10.5
// ---------------------------------------------------------
// EXERCISE 3: Convert and Fix #3
//  Fix the code.
// EXPECTED OUTPUT
//  5.5
// ---------------------------------------------------------
// EXERCISE 4: Convert and Fix #4
//  Fix the code.
// EXPECTED OUTPUT
//  9.5
// ---------------------------------------------------------
// EXERCISE: Convert and Fix #5
//  Fix the code.
// HINTS
//   maximum of int8  can be 127
//   maximum of int16 can be 32767
// EXPECTED OUTPUT
//  1127
// ---------------------------------------------------------

func main() {
	//Exercise 1
	a, b := 10, 5.5
	fmt.Println(float64(a) + b)

	//Exercise 2
	c, d := 10, 5.5
	c = int(d)
	fmt.Println(float64(c) + d)

	//Exercise 3
	// fmt.Println(int(5.5))
	fmt.Println(5.5)

	//Exercise 4
	age := 2
	fmt.Println(7.5 + float64(age))

	//Exercise 5
	// DO NOT TOUCH THESE VARIABLES
	min := int8(127)
	max := int16(1000)
	// FIX THE CODE HERE
	fmt.Println(max + int16(min))
}
