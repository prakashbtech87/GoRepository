package main

import "fmt"

func main() {

	// Println() - it is used for unformatted print, it automaticallu adds spaces between the provided argument and a new line character at the end

	// Printf() - it is used for formatted print, it allows you to control the output format by usng format specifiers
	/*
		%d: integer
		%s: String
		%T: Type of the variable
		%.3f: float
		**/

	age := 25
	name := "John Doe"
	height := 5.9

	fmt.Println("Age:", age)
	fmt.Println("Name:", name)
	fmt.Println("Height:", height)

	fmt.Printf("Age: %d, %T\n ", age, age)
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Height: %.2f\n", height)

	fmt.Printf("Age: %d, Name: %s, Height: %.2f\n", age, name, height)
}
