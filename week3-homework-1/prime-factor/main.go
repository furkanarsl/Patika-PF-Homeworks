package main

import "fmt"

func PrimeFactor(n int) []int {
	factors := make([]int, 0, 1)
	// Divide the number by 2 until we can no longer do it. Each time we divide the number add it to our factors list
	for n%2 == 0 {
		factors = append(factors, 2)
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n = n / i
		}
	}

	//When the number is a prime number
	if n > 2 {
		factors = append(factors, n)
	}
	return factors
}

func PrintFactors(factors []int) {
	for i := 0; i < len(factors); i++ {
		fmt.Printf("%v ", factors[i])
	}
	fmt.Println()
}

func main() {
	fmt.Print("Enter the number you want to calculate prime factorization: ")
	var input int
	fmt.Scanln(&input)
	PrintFactors(PrimeFactor(input))
}
