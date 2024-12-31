package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Display_Shape()
}


type Rectangle struct {
	len int
	bre int
}


func NewRect(x, y int) (*Rectangle) {
	return &Rectangle{len: x,
		bre: y,}
}

func (r *Rectangle) Display_Shape() {
	parameter := r.len * r.bre

	fmt.Println("Length is: ", r.len)
	fmt.Println("Breadth is: ", r.bre)
	fmt.Println("Parameter is: ", parameter)
}

type Circle struct {
	radius int
}

func NewCircle(r int) (*Circle) {
	return &Circle{radius: r}
}

func (c *Circle) Display_Shape() {
	parameter := 2 * math.Pi * float64(c.radius)

	fmt.Println("Radius is: ", c.radius)
	fmt.Println("Parameter is: ", parameter)	
}
