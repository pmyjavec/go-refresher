package main

import "fmt"

func main() {
	var feet float64
	var meters float64
	const diff float64 = 0.3048

	fmt.Print("Enter a measurement in feet:")
	fmt.Scanf("%f", &feet)
	meters = feet * diff
	fmt.Println(meters)
}
