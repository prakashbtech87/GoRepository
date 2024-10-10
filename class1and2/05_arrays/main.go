package main

import "fmt"

/**
Array Declaration:
- `var name [size]type`: Arrays have a fixed size, and once created, their size cannot be changed.

Array Initialization:
- Arrays can be initialized at the time of declaration or later.
  Example:
    - `var numbers [5]int`: Declares an array of integers with default values (0).
    - `var numbers = [5]int{1, 2, 3, 4, 5}`: Declares and initializes an array with values.

Accessing Array Elements:
- Elements are accessed using an index: `array[index]`.

Array Length:
- The length of an array is fixed once it is created.

In Go, arrays are initialized to their default zero values:
- Numeric types (int, float) default to `0`.
- String types default to `""` (empty string).
- Boolean types default to `false`.
- Complex types default to `nil`.
**/

func main() {
	// Declare and initialize an array to store a list of fruits
	var fruits [5]string

	// Assign values to the array
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Orange"
	fruits[3] = "Mango"
	fruits[4] = "Pineapple"

	// Access and print all array elements
	fmt.Println("List of fruits:", fruits)

	// Access and print a specific element
	fmt.Println("Third fruit:", fruits[2])

	// Initialize an array at the time of declaration
	numbers := [5]int{1, 2, 3, 4, 5}

	// Print the array elements
	fmt.Println("Numbers:", numbers)

	// Example of an array with default values (zero values for integers)
	var someNumbers [5]int
	fmt.Println("Some Numbers:", someNumbers) // Will print an array of 0s

	// Another array with string initialization
	letters := [5]string{"A", "B", "C", "D", "E"}
	fmt.Println("Letters:", letters)

	// Use %q to print the actual content of the string array in quotes
	fmt.Printf("Letters: %q\n", letters)

	// Print the length of the array
	fmt.Println("Length of letters array:", len(letters))
}
