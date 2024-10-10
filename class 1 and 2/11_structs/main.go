package main

import "fmt"

// Define a struct named person
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// create an instance of the Person struct
	// method 1: using the var keyword
	var person1 Person
	person1.FirstName = "John"
	person1.LastName = "Doe"
	person1.Age = 30

	// Access the fields of the struct
	fmt.Println("First Name: ", person1.FirstName)
	fmt.Println("Last Name: ", person1.LastName)
	fmt.Println("Age: ", person1.Age)

	// method 2: uing a struct literal
	person2 := Person{
		FirstName: "Jane",
		LastName:  "Doe",
		Age:       25,
	}
	fmt.Println("First Name: ", person2.FirstName)

	// method 3: using the new keyword (returna points to the struct)
	person3 := new(Person)
	person3.FirstName = "Alice"
	person3.LastName = "Smith"
	person3.Age = 35
}
