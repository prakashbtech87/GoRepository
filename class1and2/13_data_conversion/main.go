package main

import (
	"fmt"
	"strconv" // Import the strconv package for string conversions
)

func main() {
	// Numeric type conversion: int to float64
	var integerNum int = 40
	var floatNum float64 = float64(integerNum) // Convert int to float64
	fmt.Println("Float number:", floatNum)     // Print the float number

	// String conversion: int to string
	var number int = 50
	str := strconv.Itoa(number) // Convert integer to string
	fmt.Println("String:", str) // Print the string

	// String conversion: string to int
	strNum := "123"
	num, err := strconv.Atoi(strNum) // Convert string to integer
	fmt.Println("Number:", num)      // Print the converted number
	fmt.Println("Error:", err)       // Print any conversion error (nil if successful)

	// Convert string to float
	str2 := "3.14"
	num2, err2 := strconv.ParseFloat(str2, 64) // Convert string to float64
	if err2 == nil {                           // Check if there was no error during conversion
		fmt.Println("Float number:", num2) // Print the converted float number
	} else {
		fmt.Println("Error:", err2) // Print the error if conversion failed
	}
}
