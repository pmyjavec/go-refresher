package main

import (
	"fmt"
	"math"
)

/* What is it that defines a shape? */
type Shape struct {
}

func (s *Shape) hello() {
	fmt.Println("Hello from the land of shape")
}

/* Start Circle */
type Circle struct {
	x, y, r float64
	Shape
	Mix Shape
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return c.r*2 + math.Pi
}

/* Main */
func main() {
	c := new(Circle)
	c.r = 46.0
	c.hello()
	c.Mix.hello()
}
