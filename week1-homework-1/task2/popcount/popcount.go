package main

import "fmt"

// pc[i] is the population count of i.

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

//PopCount Using for loop for Exercise 2.3
func PopCountLoop(x uint64) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>(i*8))])
	}
	return sum
}

// PopCount using by shifting Exercise 2.4
func PopCountShift(x uint64) int {
	count := 0
	mask := uint64(1)
	for i := 0; i < 64; i++ {
		if x&mask > 0 {
			count++
		}
		x >>= 1
	}
	return count
}

// Popcount by clearing the rightmost bit Exercise 2.5
func PopCountClear(x uint64) int {
	count := 0
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}

func main() {
	count := PopCount(46)
	countLoop := PopCountLoop(46)
	countShift := PopCountShift(46)
	countClear := PopCountClear(46)
	fmt.Printf("%v\n", count)
	fmt.Printf("%v\n", countLoop)
	fmt.Printf("%v\n", countShift)
	fmt.Printf("%v\n", countClear)
}
