package main

import (
	"fmt"
	"time"
)

// ---------------------------------------------------------
// EXERCISE 1: Minutes in Weeks
//  Calculate how many minutes in two weeks.
//  STEPS:
//  1. Declare `minsPerDay` constant and initialize it
//     to the number of minutes in a day
//  2. Declare `weekDays` constant and initialize it
//     to 7.
//  3. Use printf to print the total number of minutes
//     in two weeks.
// EXPECTED OUTPUT
//  There are 20160 minutes in 2 weeks.
// ---------------------------------------------------------
// EXERCISE 2: Remove the Magic
//  Get rid of the magic numbers in the following code.
// RESTRICTIONS
//  1. You should declare 3 constants named:
//       hoursInDay, daysInWeek, and hoursPerWeek
//  2. And, hoursPerWeek constant should be initialized
//     using hoursInDay and daysInWeek constants.
// EXPECTED OUTPUT
//  There are 840 hours in 5 weeks.
// ---------------------------------------------------------
// EXERCISE 3: Constant Length
//  Calculate how many characters inside the `home`
//  constant and print it.
//  1. Declare a constant named `home`
//  2. Initialize it to "Milky Way Galaxy" string literal
//  3. Declare another constant named `length`
//  4. Initialize it by using the built-in function `len`.
//  5. Print the message below using the constants that
//     you've declared.
// RESTRICTION:
//  Use Printf.
//  Print the `home` constant using the quoted-string verb.
// EXPECTED OUTPUT
//  There are 16 characters inside "Milky Way Galaxy"
// ---------------------------------------------------------
// EXERCISE 4: TAU
//  Fix the following program and print the TAU number.
// HINT
//  You can use %g verb for printing tau.
// EXPECTED OUTPUT
//  tau = 6.283185307179586
// ---------------------------------------------------------
// EXERCISE 5: Area
//  Fix the following program.
// RESTRICTION
//  You should not use any variables.
// EXPECTED OUTPUT
//  area = 1250
// ---------------------------------------------------------
// EXERCISE 6: No Conversions Allowed
//  1. Fix the program without doing any conversion.
//  2. Explain why it doesn't work.
// EXPECTED OUTPUT
//  10h0m0s later...
// ---------------------------------------------------------
// EXERCISE 7: Iota Months
//  1. Initialize the constants using iota.
//  2. You should find the correct formula for iota.
// RESTRICTIONS
//  1. Remove the initializer values from all constants.
//  2. Then use iota once for initializing one of the
//     constants.
// EXPECTED OUTPUT
//  9 10 11
// ---------------------------------------------------------
// EXERCISE 8: Iota Months #2
//  1. Initialize multiple constants using iota.
//  2. Please follow the instructions inside the code.
// EXPECTED OUTPUT
//  1 2 3
// ---------------------------------------------------------
// EXERCISE 9: Iota Seasons
//  Use iota to initialize the season constants.
// HINT
//  You can change the order of the constants.
// EXPECTED OUTPUT
//  12 3 6 9
// ---------------------------------------------------------

func main() {
	//Exercise 1
	const (
		minsPerDay = 60 * 24
		weekDays   = 7
	)

	fmt.Printf("%d minutes in %d weeks.\n", minsPerDay*weekDays*2, 2)
	//Exercise 2
	const (
		hoursInDay   = 24
		daysInWeek   = 7
		hoursPerWeek = hoursInDay * daysInWeek
	)
	// Old version with magic numbers
	// fmt.Printf("There are %d hours in %d weeks.\n",
	// 	24*7*5, 5)
	fmt.Printf("There are %d hours in %d weeks.\n",
		hoursPerWeek*5, 5)

	//Exercise 3
	const (
		home   = "Milky Way Galaxy"
		length = len(home)
	)

	fmt.Printf("There are %d characters inside %q\n",
		length, home)

	//Exercise 4
	const (
		pi  = 3.14159265358979323846264
		tau = pi * 2
	)
	fmt.Printf("tau = %g\n", tau)

	//Exercise 5
	// const (
	// 	width  int16 = 25
	// 	height int32 = width * 2
	// )
	const (
		width  = 25
		height = width * 2
	)

	fmt.Printf("area = %d\n", width*height)
	//Exercise 6
	// We cant multiply them because time.Duration uses int64 under the hood. If we make later typless we can multiply them
	// const later int = 10
	const later = 10

	hours, _ := time.ParseDuration("1h")

	fmt.Printf("%s later...\n", hours*later)

	//Exercise 7
	const (
		Nov = 11 - iota
		Oct
		Sep
	)

	fmt.Println(Sep, Oct, Nov)

	//Exercise 8
	const (
		_ = iota
		Jan
		Feb
		Mar
	)
	// This should print: 1 2 3
	// Not: 0 1 2
	fmt.Println(Jan, Feb, Mar)

	// Exercise 9
	const (
		Spring = (iota + 1) * 3
		Summer
		Fall
		Winter
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}
