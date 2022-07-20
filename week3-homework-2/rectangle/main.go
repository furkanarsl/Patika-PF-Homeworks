package main

import (
	"errors"
	"fmt"
)

type Rectangle struct {
	length, height, area, circumference int
}

func CalcArea(r *Rectangle) {
	r.area = r.length * r.height
}

func CalcCircumference(r *Rectangle) {
	r.circumference = 2 * (r.length + r.height)
}

func CreateRectangle(length, height int) (*Rectangle, error) {
	if length < 1 || height < 1 {
		return nil, errors.New("length and height has to be greater than 0")

	}
	return &Rectangle{length: length, height: height}, nil
}

func main() {
	rect, err := CreateRectangle(10, 12)
	if err != nil {
		fmt.Println(err)
	}
	CalcArea(rect)
	CalcCircumference(rect)
	fmt.Printf("Rectangle length: %v height: %v area: %v circumference: %v\n", rect.height, rect.length, rect.area, rect.circumference)
}
