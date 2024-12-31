package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello the king in black")
	r1 := NewRect(5, 6)

	c1 := NewCircle(10)

	r1.Display_Shape()
	c1.Display_Shape()
}
