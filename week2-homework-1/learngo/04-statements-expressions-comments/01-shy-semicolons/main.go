package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Shy Semicolons
//
//  1. Try to type your statements by separating them using
//     semicolons
//
//  2. Observe how Go fixes them for you
//
// ---------------------------------------------------------

func main() {
	// Before formatting the code
	// fmt.Println("test");fmt.Println("test")

	// After formatting the code gofmt removes the semicolons
	fmt.Println("test")
	fmt.Println("test")
}
