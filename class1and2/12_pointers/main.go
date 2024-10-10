package main

import "fmt"

func main() {
	// Declare an integer variable
	var num int = 2

	// Declare a pointer to an integer
	var ptr *int

	// Assign the address of num to ptr
	ptr = &num

	// Print the value of the variable num
	fmt.Println("Value of variable:", num)
	// Print the value of the variable using the pointer
	fmt.Println("Value via pointer:", *ptr)

	// Declare a string variable
	data := "Go Lang"
	// Declare a pointer to the string
	pointer := &data
	// Print the value of the string using the pointer
	fmt.Println("Value via pointer:", *pointer)

	// Declare another pointer to an integer
	var ptr2 *int
	// Check if ptr2 is nil (not initialized)
	if ptr2 == nil {
		fmt.Println("Nil pointer") // This will be printed
	}
}
