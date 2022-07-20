package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE 1: Windows Path
//  1. Change the following program
//  2. It should use a raw string literal instead
// HINT
//  Run this program first to see its output.
//  Then you can easily understand what it does.
// EXPECTED OUTPUT
//  Your solution should output the same as this program.
//  Only that it should use a raw string literal instead.
// ---------------------------------------------------------
// EXERCISE 2: Print JSON
//  1. Change the following program
//  2. It should use a raw string literal instead
// HINT
//  Run this program first to see its output.
//  Then you can easily understand what it does.
// EXPECTED OUTPUT
//  Your solution should output the same as this program.
//  Only that it should use a raw string literal instead.
// ---------------------------------------------------------
// EXERCISE 3: Raw Concat
//  1. Initialize the name variable
//     by getting input from the command line
//  2. Use concatenation operator to concatenate it
//     with the raw string literal below
// NOTE
//  You should concatenate the name variable in the correct
//  place.
// EXPECTED OUTPUT
//  Let's say that you run the program like this:
//   go run main.go inanç
//  Then it should output this:
//   hi inanç!
//   how are you?
// ---------------------------------------------------------
// EXERCISE 4: Count the Chars
//  1. Change the following program to work with unicode
//     characters.
// INPUT
//  "İNANÇ"
// EXPECTED OUTPUT
//  5
// ---------------------------------------------------------
// EXERCISE 5: Improved Banger
//  Change the Banger program the work with Unicode
//  characters.
// INPUT
//  "İNANÇ"
// EXPECTED OUTPUT
//  İNANÇ!!!!!
// ---------------------------------------------------------
// EXERCISE 6: ToLowercase
//  1. Look at the documentation of strings package
//  2. Find a function that changes the letters into lowercase
//  3. Get a value from the command-line
//  4. Print the given value in lowercase letters
// HINT
//  Check out the strings package from Go online documentation.
//  You will find the lowercase function there.
// INPUT
//  "SHEPARD"
// EXPECTED OUTPUT
//  shepard
// ---------------------------------------------------------
// EXERCISE 7: Trim It
//  1. Look at the documentation of strings package
//  2. Find a function that trims the spaces from
//     the given string
//  3. Trim the text variable and print it
// EXPECTED OUTPUT
//  The weather looks good.
//  I should go and play.
// ---------------------------------------------------------
// EXERCISE 8: Right Trim It
//  1. Look at the documentation of strings package
//  2. Find a function that trims the spaces from
//     only the right-most part of the given string
//  3. Trim it from the right part only
//  4. Print the number of characters it contains.
// RESTRICTION
//  Your program should work with unicode string values.
// EXPECTED OUTPUT
//  5
// ---------------------------------------------------------

func main() {
	//Exercise 1
	// path := "c:\\program files\\duper super\\fun.txt\n" +
	// 		"c:\\program files\\really\\funny.png"
	path := `c:\program files\duper super\fun.txt
c:\program files\really\funny.png`
	fmt.Println(path)

	//Exercise 2
	// json := "\n" +
	// 	"{\n" +
	// 	"\t\"Items\": [{\n" +
	// 	"\t\t\"Item\": {\n" +
	// 	"\t\t\t\"name\": \"Teddy Bear\"\n" +
	// 	"\t\t}\n" +
	// 	"\t}]\n" +
	// 	"}\n"

	json := `
{
	"Items": [{
		"Item": {
			"name": "Teddy Bear"
		}
	}]
}`
	fmt.Println(json)

	//Exercise 3
	name := os.Args[1]

	msg := `hi ` + name + `!
how are you?`

	fmt.Println(msg)

	//Exercise 4
	// length := len(os.Args[1])
	length := utf8.RuneCountInString(os.Args[1])
	fmt.Println(length)

	//Exercise 5
	msg2 := os.Args[1]

	// s := msg2 + strings.Repeat("!", len(msg2))
	l := utf8.RuneCountInString(msg2)
	s := msg2 + strings.Repeat("!", l)

	fmt.Println(s)

	//Exercise 6
	//msg2 is from the previous exercise.
	fmt.Println(strings.ToLower(msg2))

	//Exercise 7
	m := `
	
	The weather looks good.
I should go and play.
	`

	fmt.Println(strings.TrimSpace(m))

	//Exercise 8
	n := "inanç           "
	n = strings.TrimRight(n, " ")
	l2 := utf8.RuneCountInString(n)
	fmt.Println(l2)
}
