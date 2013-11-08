package main

import "fmt"

func main() {
	// var x string = "Hello World"
	// fmt.Println(x)

	var input float64

	// fmt.Print("Enter a number: ")
	// fmt.Scanf("%f", &input)

	// output := input * 2
	// fmt.Println(output)

	// fmt.Print("Enter degrees in F: ")
	// fmt.Scanf("%f", &input)
	// fmt.Println(input, "degrees F is", farenheitToCelcius(input), "degrees C")

	fmt.Print("Enter feet: ")
	fmt.Scanf("%f", &input)
	fmt.Println(input, "feet is", feetToMeters(input), "meters")
}

func farenheitToCelcius(degreesF float64) float64 {
	return (degreesF - 32) * 5 / 9
}

func feetToMeters(feet float64) float64 {
	return feet * 0.3048
}
