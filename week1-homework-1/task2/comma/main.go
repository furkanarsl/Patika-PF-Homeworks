package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// Comma without using recursion exercise 3.10
func commaNoRecursion(s string) string {
	b := &bytes.Buffer{}
	pre := len(s) % 3

	if pre == 0 {
		pre = 3
	}

	b.WriteString(s[:pre])
	// Deal with the rest.
	for i := pre; i < len(s); i += 3 {
		b.WriteByte(',')
		b.WriteString(s[i : i+3])
	}
	return b.String()
}
func main() {
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
	fmt.Println(commaNoRecursion("1234567"))
}
