package main

import "fmt"

// Define a struct named Person
type Person struct {
	FirstName string // First name of the person
	LastName  string // Last name of the person
	Age       int    // Age of the person
}

func main() {
	// Method 1: Creating an instance using the var keyword
	var person1 Person
	person1.FirstName = "John" // Set the first name
	person1.LastName = "Doe"   // Set the last name
	person1.Age = 30           // Set the age

	// Access and print the fields of person1
	fmt.Println("Person 1:")
	fmt.Println("First Name:", person1.FirstName)
	fmt.Println("Last Name:", person1.LastName)
	fmt.Println("Age:", person1.Age)

	// Method 2: Creating an instance using a struct literal
	person2 := Person{
		FirstName: "Jane", // Set the first name
		LastName:  "Doe",  // Set the last name
		Age:       25,     // Set the age
	}

	// Access and print the first name of person2
	fmt.Println("\nPerson 2:")
	fmt.Println("First Name:", person2.FirstName)

	// Method 3: Creating an instance using the new keyword
	person3 := new(Person)      // Returns a pointer to the struct
	person3.FirstName = "Alice" // Set the first name
	person3.LastName = "Smith"  // Set the last name
	person3.Age = 35            // Set the age

	// Access and print the details of person3
	fmt.Println("\nPerson 3:")
	fmt.Println("First Name:", person3.FirstName)
	fmt.Println("Last Name:", person3.LastName)
	fmt.Println("Age:", person3.Age)
}
