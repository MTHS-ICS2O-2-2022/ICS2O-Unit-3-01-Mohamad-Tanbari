// Made by Mohamad
// Made on Mar 30 2023
//
// This program calculates the area of a trapezoid

package main

import (
	"fmt"
)

func main() {
	// Declare variables
	var height float64
	var aBase float64
	var bBase float64

	// Get user input
	fmt.Print("Enter the height of the trapezoid: ")
	fmt.Scanln(&height)
	fmt.Print("Enter the length of the first base: ")
	fmt.Scanln(&aBase)
	fmt.Print("Enter the length of the second base: ")
	fmt.Scanln(&bBase)

	// Calculate the area
	area := (aBase + bBase) / 2 * height

	// Output area
	fmt.Println("\nThe area of the trapezoid is", area, "cm²")

	fmt.Println("\nDone.")
}
