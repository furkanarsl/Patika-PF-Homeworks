package main

import "fmt"

//Exercise 10
var isValid bool

// ---------------------------------------------------------
// EXERCISE: Declare int
//
//  1. Declare and print a variable with an int type
//  2. The declared variable's name should be: height
// EXPECTED OUTPUT
//  0
// ---------------------------------------------------------
// EXERCISE 2: Declare bool
//  1. Declare and print a bool variable
//  2. The variable's name should be: isOn
// EXPECTED OUTPUT
//  false
// ---------------------------------------------------------
// EXERCISE 3: Declare float64
//  1. Declare and print a variable with a float64 type
//  2. The declared variable's name should be: brightness
// EXPECTED OUTPUT
//  0
// ---------------------------------------------------------
// EXERCISE 4: Declare string
//  1. Declare a string variable
//  2. Print that variable
//
// EXPECTED OUTPUT
//  ""
// ---------------------------------------------------------
// EXERCISE 5: Undeclarables
//  1. Declare the variables below:
//      3speed
//      !speed
//      spe?ed
//      var
//      func
//      package
//  2. Observe the error messages
// NOTE
//  The types of the variables are not important.
// ---------------------------------------------------------
// EXERCISE 6: Declare with bits
//
//  1. Declare a few variables using the following types
//    int
//    int8
//    int16
//    int32
//    int64
//    float32
//    float64
//    complex64
//    complex128
//    bool
//    string
//    rune
//    byte
//
// 2. Observe their output
// 3. After you've done, check out the solution
//    and read the comments there
// EXPECTED OUTPUT
//  0 0 0 0 0 0 0 (0+0i) (0+0i) false 0 0
//  ""
// EXERCISE 7: Multiple
//
//  1. Declare two variables using
//     multiple variable declaration statement
//  2. The first variable's name should be active
//  3. The second variable's name should be delta
//  4. Print them all
// HINT
//  You should declare a bool and an int variable
//
// EXPECTED OUTPUT
//  false 0
// ---------------------------------------------------------
// EXERCISE 8: Multiple #2
//
//  1. Declare and initialize two string variables
//     using multiple variable declaration
//
//  2. Use the type once while declaring the variables
//
//  3. The first variable's name should be firstName
//  4. The second variable's name should be lastName
//
//  5. Print them all
//
// EXPECTED OUTPUT
//  "" ""
// ---------------------------------------------------------
// EXERCISE 9: Unused
//
//  1. Declare a variable
//  2. Variable's name should be: isLiquid
//  3. Discard it using a blank-identifier
//
// NOTE
//  Do not print the variable
// ---------------------------------------------------------
// EXERCISE 10: Package Variable
//  1. Declare a variable in the package-scope
//  2. Observe whether something happens when you don't
//     use it
// ---------------------------------------------------------
// EXERCISE 11: Wrong doer
//  1. Print a variable
//  2. Then declare it
//  (This means: Try to print it before its declaration)
//  3. Observe the error
// ---------------------------------------------------------
func main() {
	var height int
	fmt.Println(height)

	// Exercise 2
	var isOn bool
	fmt.Println(isOn)

	//Exercise 3
	var brightness float64
	fmt.Println(brightness)

	//Exercise 4
	var name string
	fmt.Printf("s (%T): %q\n", name, name)

	//Exercise 5
	// Commented out because of error messages.
	// var 3speed int
	// var !speed int
	// var spe?ed int
	// var var int
	// var func int
	// var package int

	// Exercise 6
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	var f32 float32
	var f64 float64

	var c64 complex64
	var c128 complex128

	var b bool

	var s string
	var r rune
	var by byte
	fmt.Println(i, i8, i16, i32, i64, f32, f64, c64, c128, b, s, r, by)

	// Exercise 7
	var (
		active bool
		delta  int
	)
	fmt.Println(active, delta)

	//Exercise 8
	var firstName, lastName string = "", ""
	fmt.Println(firstName, lastName)

	//Exercise 9
	var isLiquid bool
	_ = isLiquid

	// Exercise 11
	// Printing before declaring the variable gives error.
	// fmt.Println(username)
	// var username string
}
