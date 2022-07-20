package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE 1: Count the Arguments
//  Print the count of the command-line arguments
// INPUT
//  bilbo balbo bungo
// EXPECTED OUTPUT
//  There are 3 names.
// ---------------------------------------------------------
// EXERCISE 2: Print the Path
//  Print the path of the running program by getting it
//  from `os.Args` variable.
// HINT
//  Use `go build` to build your program.
//  Then run it using the compiled executable program file.
// EXPECTED OUTPUT SHOULD INCLUDE THIS
//  myprogram
// ---------------------------------------------------------
// EXERCISE 3: Print Your Name
//  Get it from the first command-line argument
// INPUT
//  Call the program using your name
// EXPECTED OUTPUT
//  It should print your name
// EXAMPLE
//  go run main.go inanc
//    inanc
// BONUS: Make the output like this:
//  go run main.go inanc
//    Hi inanc
//    How are you?
// ---------------------------------------------------------
// EXERCISE 4: Greet More People
// RESTRICTIONS
//  1. Be sure to match the expected output below
//  2. Store the length of the arguments in a variable
//  3. Store all the arguments in variables as well
// INPUT
//  bilbo balbo bungo
// EXPECTED OUTPUT
//  There are 3 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Nice to meet you all.
// ---------------------------------------------------------
// EXERCISE 5: Greet 5 People
//  Greet 5 people this time.
//  Please do not copy paste from the previous exercise!
// RESTRICTION
//  This time do not use variables.
// INPUT
//  bilbo balbo bungo gandalf legolas
// EXPECTED OUTPUT
//  There are 5 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Hello great gandalf !
//  Hello great legolas !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	// Since I combined all the exercises into one file and Exercise#5 requires 5 command line arguments
	// you have to provide 5 of them for this program or there will be an error.

	// Exercise 1
	// UNCOMMENT & FIX THIS CODE
	count := len(os.Args) - 1

	// UNCOMMENT IT & THEN DO NOT TOUCH THIS CODE
	fmt.Printf("There are %d names.\n", count)

	// Exercise 2
	// go build -o myprogram
	fmt.Println(os.Args[0])

	//Exercise 3
	// go run main.go yourNameHere
	fmt.Println("Hi", os.Args[1])
	fmt.Println("How are you?")

	//Exercise 4
	var (
		nPeople = len(os.Args) - 1
		p1      = os.Args[1]
		p2      = os.Args[2]
		p3      = os.Args[3]
	)

	fmt.Println("There are", nPeople, "people !")
	fmt.Println("Hello great", p1, "!")
	fmt.Println("Hello great", p2, "!")
	fmt.Println("Hello great", p3, "!")
	fmt.Println("Nice to meet you all.")

	// Exercise 5
	fmt.Println("There are", len(os.Args)-1, "people !")
	fmt.Println("Hello great", os.Args[1], "!")
	fmt.Println("Hello great", os.Args[2], "!")
	fmt.Println("Hello great", os.Args[3], "!")
	fmt.Println("Hello great", os.Args[4], "!")
	fmt.Println("Hello great", os.Args[5], "!")
	fmt.Println("Nice to meet you all.")
}
