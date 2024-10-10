package main

import "fmt"

// Function without parameters and return type
func simpleFunction() {
	fmt.Println("Simple Function")
}

// Function with parameters and return type
func add(a int, b int) int {
	// Returns the sum of a and b
	return a + b
}

// Function with multiple parameters
func add2(a, b, c, d, e, f, g int) int {
	// Adds seven integers and returns the sum
	return a + b + c + d + e + f + g
}

// Function with a named return variable
func multiply(a, b int) (result int) {
	// Multiplies a and b and returns the result
	result = a * b
	return // Implicit return of 'result'
}

// Function that returns a value and an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		// If division by zero, return an error
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	// Otherwise, return the result and nil error
	return a / b, nil
}

func main() {
	// Calling a function without parameters
	simpleFunction()

	// Calling a function with two parameters
	sum := add(10, 20)
	fmt.Println("Sum:", sum)

	// Calling a function with multiple parameters
	sum2 := add2(10, 20, 30, 40, 50, 60, 70)
	fmt.Println("Sum2:", sum2)

	// Calling a function with a named return value
	multiplyResult := multiply(10, 20)
	fmt.Println("Multiply:", multiplyResult)

	// Calling a function that returns a value and an error
	divideResult, _ := divide(10, 5) // Ignoring the error for simplicity
	fmt.Println("Divide:", divideResult)

	// Example of handling an error
	// divideResult, err := divide(10, 0)
	// if err != nil {
	//     fmt.Println("Error:", err)
	// } else {
	//     fmt.Println("Divide:", divideResult)
	// }
}
