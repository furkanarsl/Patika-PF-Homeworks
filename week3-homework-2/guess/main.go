package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func generateNumber() []int {
	d := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Seed(time.Now().UnixNano())
	for d[0] == 0 {
		rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
	}
	n := d[:4]
	return n
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func checkGuess(number []int, guess int) string {
	correctPos, correctNum := 0, 0
	guessStr := strconv.Itoa(guess)

	for i := 0; i < 4; i++ {
		digit, _ := strconv.Atoi(string(guessStr[i]))
		if contains(number, digit) {
			if digit == number[i] {
				correctPos += 1
			} else {
				correctNum += 1
			}
		}
	}
	return fmt.Sprintf("+%v-%v", correctPos, correctNum)

}

func main() {
	min, max := 1000, 9999
	prevResults := make(map[int]string)
	n := generateNumber()
	// fmt.Println(n)
	var result string
	var guess int
	for result != "+4-0" {
		fmt.Print("Enter your guess: \n")
		fmt.Scanln(&guess)
		if guess < min || guess > max {
			fmt.Println("Your guess is not in the range. The number contains 4 digits. Please try again.")
			continue
		}
		// fmt.Print("\033[H\033[2J")
		result = checkGuess(n, guess)
		fmt.Println("**************************")
		fmt.Println("Your previous guesses: ")
		for k, v := range prevResults {
			fmt.Printf("Guess: %v Result: %v\n", k, v)
		}
		fmt.Printf("Latests Guess: %v Result: %v\n", guess, result)
		prevResults[guess] = result

		fmt.Println("**************************")
	}
	fmt.Println("Congrats you found the number!")
}
