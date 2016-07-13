package main

import (
	"fmt"
	"math"
)

/* What is it that defines a shape? */
type Shape interface {
	area() float64
	perimeter() float64
}

/* Start Circle */
type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return c.r*2 + math.Pi
}

/* Start Rectangle */
type Rectangle struct {
	x, y float64
}

func (r *Rectangle) area() float64 {
	return r.x * r.y
}

func (r *Rectangle) perimeter() float64 {
	return r.x*2 + r.y*2
}

/* Print areas of shapes */
func areaPrinter(shapes ...Shape) {
	for _, s := range shapes {
		fmt.Println(s.area())
	}
}

/* Main */
func main() {
	c := new(Circle)
	c.r = 46.0

	r := Rectangle{x: 10, y: 15}

	areaPrinter(c, &r)
}
