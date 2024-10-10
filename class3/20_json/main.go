package main

import (
	"encoding/json" // Import the encoding/json package for JSON encoding/decoding
	"fmt"           // Import the fmt package for formatted I/O
)

// Define a struct named Person with JSON field tags
type Person struct {
	Name    string `json:"name"`     // The Name field will be encoded as "name"
	Age     int    `json:"age"`      // The Age field will be encoded as "age"
	IsAdult bool   `json:"is_adult"` // The IsAdult field will be encoded as "is_adult"
}

func main() {
	// Encoding
	person := Person{Name: "Jane Doe", Age: 25, IsAdult: true} // Create a new Person instance
	jsonData, err := json.Marshal(person)                      // Convert the struct to JSON
	if err != nil {
		fmt.Println("Error encoding JSON:", err) // Handle encoding error
		return
	}
	fmt.Println("JSON data:", string(jsonData)) // Print the JSON data

	// Decoding
	var newPerson Person                       // Declare a new Person variable for decoding
	err = json.Unmarshal(jsonData, &newPerson) // Convert JSON data back into the struct
	if err != nil {
		fmt.Println("Error decoding JSON:", err) // Handle decoding error
		return
	}

	fmt.Println("Decoded Person:", newPerson) // Print the decoded Person instance
}
