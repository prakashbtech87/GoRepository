package main

import "fmt"

func main() {

	// Println() - Used for unformatted printing
	// It automatically adds spaces between the arguments and ends with a newline
	age := 25
	name := "John Doe"
	height := 5.9

	fmt.Println("Age:", age)       // Prints age with automatic formatting
	fmt.Println("Name:", name)     // Prints name
	fmt.Println("Height:", height) // Prints height

	// Printf() - Used for formatted printing
	// Format specifiers:
	// %d: integer
	// %s: string
	// %T: type of the variable
	// %.2f: float (formatted to 2 decimal places)

	fmt.Printf("Age: %d, %T\n", age, age) // Prints age with type
	fmt.Printf("Name: %s\n", name)        // Prints name as a string
	fmt.Printf("Height: %.2f\n", height)  // Prints height as a float with 2 decimal places

	// Combining multiple format specifiers in one Printf statement
	fmt.Printf("Age: %d, Name: %s, Height: %.2f\n", age, name, height)
}
