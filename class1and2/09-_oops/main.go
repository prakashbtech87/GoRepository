package main

import "fmt"

func main() {
	// Example 1: Simple for loop
	// This loop prints numbers from 0 to 4
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Example 2: Infinite loop with a break condition
	counter := 0
	for {
		// Print message for infinite loop
		fmt.Println("Infinite loop")
		counter++ // Increment counter
		// Break the loop if counter exceeds 5
		if counter > 5 {
			break // Exit the infinite loop
		}
	}

	// Example 3: Looping through a slice of integers
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		// Print index and value from the slice
		fmt.Println("\nIndex:", index, "Value:", value)
	}

	// Example 4: Looping through a string
	message := "Hello, World!"
	for index, char := range message {
		// Print index and character from the string
		fmt.Printf("\nIndex: %d, Character: %c\n", index, char)
	}
}
