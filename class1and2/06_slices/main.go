package main

import "fmt"

func main() {

	// Declare and initialize a slice of integers with values
	numbers := []int{1, 2, 3, 4, 5}

	// Access and print slice elements
	fmt.Println("Slice elements: ", numbers)

	// Print the length of the slice
	fmt.Println("Length of the slice: ", len(numbers))

	// Create another slice without specifying capacity
	nums := []int{1, 2, 3, 4}

	// Print slice elements, its length, and capacity
	fmt.Println("Slice elements:", nums)
	fmt.Println("Length:", len(nums))   // The number of elements
	fmt.Println("Capacity:", cap(nums)) // The capacity of the underlying array

	// Creating a slice using the `make` function with a length of 3 and a capacity of 5
	numsCheck := make([]int, 3, 5)

	// Appending elements to the slice (this will increase its size)
	numsCheck = append(numsCheck, 4) // Adding element 4
	numsCheck = append(numsCheck, 5) // Adding element 5
	numsCheck = append(numsCheck, 6) // Adding element 6

	// Appending more elements to test slice growth
	numsCheck = append(numsCheck, 4)
	numsCheck = append(numsCheck, 5)
	numsCheck = append(numsCheck, 6)
	numsCheck = append(numsCheck, 4)
	numsCheck = append(numsCheck, 5)

	// Print the modified slice and its updated length and capacity
	fmt.Println("\n\nSliced array:", numsCheck)                         // Output the full slice
	fmt.Println("Length:", len(numsCheck), "Capacity:", cap(numsCheck)) // Print length and capacity
}
