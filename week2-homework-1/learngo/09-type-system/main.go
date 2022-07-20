package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE 1: Optimal Types
//
//  1. Choose the optimal data types for the given situations.
//  2. Print them all
//  3. Try converting them to lesser data types.
//     For example try converting int64 variable to int32.
//     Then observe the result.
//     Search the web why the result is so?
//
// NOTE
//  This is just an exercise for teaching you different
//  data types. Do not apply it to the real-life.
//
//  As I said in the lectures that, premature optimization
//  is not a good thing.
// ---------------------------------------------------------
// EXERCISE: The Type Problem
//  Solve the data type problem in the program.
// EXPECTED OUTPUT
//  width: 265 height: 265
//  are they equal? true
// ---------------------------------------------------------

func main() {
	//Exercise 1
	// an english letter (search web for: ascii control code)
	var letter byte = 'f'
	fmt.Println(letter)
	// a non-english letter (search web for: unicode codepoint)
	var unicode rune = 'Ğ'
	fmt.Println(unicode)
	// a year in 4 digits like 2040
	var year uint16 = 2040
	fmt.Println(year)
	// a month in 2 digits: 1 to 12
	var month uint8 = 4
	fmt.Println(month)
	// the speed of the light
	var speedOfLight uint64 = 299792458
	fmt.Println(speedOfLight)
	// angle of a circle
	var angle float32 = 16.5
	fmt.Println(angle)
	// PI number: 3.141592653589793
	var pi float64 = 3.141592653589793
	fmt.Println(pi)

	//Exercise 2
	var (
		// width  uint8
		width  uint16
		height uint16
	)
	width, height = 255, 265
	width += 10
	fmt.Printf("width: %d height: %d\n", width, height)

	fmt.Println("are they equal?", width == height)
}
