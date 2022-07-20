package main

import (
	"fmt"
	"strings"
)

// Exercise 3.12 Anagram function
func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	a = strings.ToLower(a)
	b = strings.ToLower(b)

	aLetters := make(map[rune]int)
	// Count the letters in a
	for _, c := range a {
		aLetters[c]++
	}

	// Counts letters of b
	for _, c := range b {
		// If a letter of b does not exists in a we can return false early
		if _, ok := aLetters[c]; !ok {
			return false
		} else {
			aLetters[c]--
			// If b has more of the same letter we can also return false early
			if aLetters[c] < 0 {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("secure", "rescue"))
	fmt.Println(isAnagram("Bernard Lowe", "Arnold Weber"))
}
