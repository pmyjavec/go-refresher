package math

import (
	"fmt"
)

func max(ints ...int) int {
	biggest := ints[0]

	if len(ints) == 1 {
		return biggest
	}

	for _, v := range ints {
		if v > biggest {
			biggest = v
		}
	}

	return biggest
}

func main() {
	biggest := max()
	fmt.Println(biggest)
}
