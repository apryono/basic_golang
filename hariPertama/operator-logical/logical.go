/*
	Ini Merupakan Materi Operator Artimetika
*/
package main

import "fmt"

func main() {
	// Logical Operator

	condLeft := true
	condRight := false

	// condResult := condLeft && condRight // logical and
	// condResult := condLeft || condRight // logical or
	condResult := !condLeft || condRight // logical use negasi

	fmt.Println("condLeft :", condLeft)
	fmt.Println("condRight :", condRight)
	fmt.Println("condResult :", condResult)
}
