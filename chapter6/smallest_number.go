package main

import (
	"fmt"
)

func main() {
	numbers := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	smallest := 0
	for k, v := range numbers {

		if k == 0 {
			smallest = v
		}

		if v < smallest {
			smallest = v
		}
	}

	fmt.Println("Smallest number is:", smallest)
}
