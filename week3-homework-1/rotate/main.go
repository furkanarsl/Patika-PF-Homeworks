package main

import "fmt"

func Rotate(input []int, steps int, rotateLeft bool) []int {
	if steps < 0 || len(input) == 0 {
		return input
	}
	k := steps % len(input)
	if rotateLeft {
		input = append(input[k:], input[:k]...)
	} else {
		diff := len(input) - k
		input = append(input[diff:], input[:diff]...)
	}
	return input
}

func main() {
	n := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("Before rotation: %v\n", n)
	n = Rotate(n, 3, true)
	fmt.Printf("After rotation:  %v\n", n)
}
