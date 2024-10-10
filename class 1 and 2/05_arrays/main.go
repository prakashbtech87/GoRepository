package main

/**
Array Declaration - var name [size]type
array are fixed size, once the array is created its size cannot be changed

Array Initialization - can be initiualized at the time of declaration and later
var numbers [5]int - create an array and initialized with 0
var numbers = [5]int{1, 2, 3, 4, 5} - create an array and initialized with 1, 2, 3, 4, 5

Accessing array elements - array elements are accessed using index with []


Array length - length canot be changed once the array is created

in Go when you declare an array the elements are initlaize to their zero value (default value)
numeric (int, float) - 0
string - ""
boolean - false
complex - nil

**/

import "fmt"

func main() {
	// declare and initialize array to store a list of fruits
	var fruits [5]string

	// asssign values to the array
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Orange"
	fruits[3] = "Mango"
	fruits[4] = "Pineapple"

	// access and print array element
	fmt.Println("List of fruits:", fruits)

	// access and print a specific element
	fmt.Println("First fruit:", fruits[2])

	// initialize array at the time of declaration in a single line
	numbers := [5]int{1, 2, 3, 4, 5}

	// access and print array elements
	fmt.Println("Numbers:", numbers)

	var someNumbers [5]int
	fmt.Println("Some Numbers:", someNumbers)

	letters := [5]string{"A", "B", "C", "D", "E"}
	fmt.Println("Letters:", letters)

	// if you wan to see the actual content of the array
	fmt.Printf("Letters: %q\n", letters)

	// print the length of the array
	fmt.Println("Length of letters array:", len(letters))

}
