package main

import "fmt"

func main() {
	const (
		January = 1 + iota
		February
		March
		April
		May
		June
		July
		August
		September
		October
		November
		December
	)

	fmt.Println("Current month is:", July)
}
