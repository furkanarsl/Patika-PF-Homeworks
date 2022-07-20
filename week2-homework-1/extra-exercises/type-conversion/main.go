package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i64 int64 = 12
	var f64 float64 = 1234.273

	fmt.Println("f64 as float32: ", float32(f64))
	fmt.Println("f64 as int64: ", int64(f64))
	fmt.Println("i64 as float64: ", float64(i64))

	a := "123456"
	b := "435.16"
	c, _ := strconv.ParseInt(a, 10, 64)
	d, _ := strconv.ParseFloat(b, 64)

	fmt.Println("var a converted to an integer: ", c)
	fmt.Println("var d converted to an integer: ", d)

}
