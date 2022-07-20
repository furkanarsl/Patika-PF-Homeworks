package main

import "fmt"

func Sieve(n int) []bool {
	// Create array for integers from 0 to n
	// If ith position is true the number is prime.
	numbers := make([]bool, n+1)

	// Set all values starting from 2 to True
	for i := 2; i < n+1; i++ {
		numbers[i] = true
	}

	//Starting at 2 we check if the number is prime. Since 2 is prime we start marking multiples of 2 as non primes.
	for p := 2; p*p <= n; p++ {
		if numbers[p] {
			for i := p * 2; i <= n; i += p {
				numbers[i] = false
			}
		}
	}
	return numbers
}

func PrintSieve(sieve []bool) {
	for i := 0; i < len(sieve); i++ {
		if sieve[i] {
			fmt.Printf("%v ", i)
		}
	}
	fmt.Println()
}
func main() {
	fmt.Print("Enter the number you want to find the primes up to: ")
	var input int
	fmt.Scanln(&input)
	PrintSieve(Sieve(input))
}
