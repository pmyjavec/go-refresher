package main

import "fmt"

func main() {
	fmt.Print("Enter a temperature (F) to convert to celcius (C): ")
	var fahrenheit float64

	fmt.Scanf("%f", &fahrenheit)

	C := (fahrenheit - 32) * 5 / 9

	fmt.Println(C)
}
