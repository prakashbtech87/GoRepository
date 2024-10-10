package main

import "fmt"

func main() {
	// create a map
	//studentGrades := make(map[string]int)

	// adding key-value pairs
	/*studentGrades["John"] = 92
	studentGrades["Jane"] = 85
	studentGrades["Doe"] = 88*/

	studentGrades := map[string]int{
		"John":  92,
		"Jane":  85,
		"Doe":   88,
		"Alice": 78,
	}

	// accessing value
	fmt.Println("John's grade: ", studentGrades["John"])

	// modifying value
	studentGrades["John"] = 100
	fmt.Println("John's grade: ", studentGrades["John"])

	// deleting a key-value pair
	delete(studentGrades, "Jane")

	// checking if a key exists
	grade, exists := studentGrades["Jane"]
	fmt.Println("Jane's grade: ", grade, "Exists: ", exists)

	// iterating over a map
	fmt.Println("Student grades:")
	for name, grade := range studentGrades {
		fmt.Println(name, ":", grade)
	}

}
