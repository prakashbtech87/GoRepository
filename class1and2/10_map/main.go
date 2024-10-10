package main

import "fmt"

func main() {
	// Create a map to store student names and their grades
	studentGrades := make(map[string]int)

	// Adding key-value pairs (student name and their grade)
	studentGrades["John"] = 92
	studentGrades["Jane"] = 85
	studentGrades["Doe"] = 88

	// Accessing a value from the map using a key
	fmt.Println("John's grade: ", studentGrades["John"])

	// Modifying an existing value
	studentGrades["John"] = 100
	fmt.Println("John's updated grade: ", studentGrades["John"])

	// Deleting a key-value pair (removing Jane's grade)
	delete(studentGrades, "Jane")

	// Checking if a key exists in the map
	grade, exists := studentGrades["Jane"]
	fmt.Println("Jane's grade:", grade, "Exists:", exists) // 'exists' is false if key not found

	// Iterating over the map to print all student grades
	fmt.Println("Student grades:")
	for name, grade := range studentGrades {
		fmt.Println(name, ":", grade)
	}
}
