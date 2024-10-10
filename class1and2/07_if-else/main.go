package main

import "fmt"

func main() {
	// Initialize variable x with value 10
	x := 10

	// Check if x is greater than 5
	if x > 5 {
		// If true, print this message
		fmt.Println("x is greater than 5")
	} else {
		// If false, print this message
		fmt.Println("x is not greater than 5")
	}

	// Initialize variable z with value 7
	z := 7

	// Check if z is greater than 10
	if z > 10 {
		// If true, print this message
		fmt.Println("z is greater than 10")
	} else if z > 5 {
		// Check if z is greater than 5 but not greater than 10
		fmt.Println("z is greater than 5")
	} else {
		// If none of the conditions match, z is less than or equal to 5
		fmt.Println("z is less than or equal to 5")
	}
}
