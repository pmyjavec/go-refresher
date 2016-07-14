package main

import (
	"fmt"
)

func fibGenerator() func() uint {
	x := uint(1)
	y := uint(0)
	result := uint(0)
	return func() uint {
		result = x + y
		y = x
		x = result
		return x
	}
}

func main() {
	fib := fibGenerator()
	for i := 0; i < 100; i++ {
		fmt.Println(fib())
	}
}
