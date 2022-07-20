package main

import "fmt"

func main() {
	x, y, z := 10, 20, 30
	fmt.Println(x * z)
	fmt.Println(z / x)
	fmt.Println(y + x)
	fmt.Println(y - x)
	fmt.Println(z % x)
	x--
	z++
	fmt.Println(x)
	fmt.Println(z)
}
