package main

import (
	"fmt"
)

func half(num int) (float64, bool) {
	halved := (float64(num) / 2)
	parity := num%2 == 0
	return halved, parity
}

func main() {
	for i := 0; i < 10; i++ {
		x, y := half(i)
		fmt.Println(i, ":", x, y)
	}
}
