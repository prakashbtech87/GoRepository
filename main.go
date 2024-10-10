package main

import "fmt"

func main() {
	// variables with different data types
	// var name = "John Doe"
	// var age = 25

	var name string = "John Doe"
	var age int = 25
	var height float64 = 5.9
	var isMarried bool = true

	isMarried = false

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Married:", isMarried)

	const pi float64 = 3.14
	//pi = 6.14

	fmt.Println("PI:", pi)

	//var rollNumber int = 101
	// Shorthand variable initalization
	rollNumber := 101
	fmt.Println("Roll Number:", rollNumber)

	// in Go, the visibility of a variable or function outside its package is determined by the capitalization of its name
	// if the first letter of a name is uppercase, its exported (public)
	// if the first letter of a name is lowercase, its private and only visible within same package

	// public variable
	var PublicVariable int = 100

	// Private variable
	var privateVariable int = 200

	fmt.Println("Public Variable:", PublicVariable)
	fmt.Println("Private Variable:", privateVariable)

}
