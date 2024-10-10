package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	// Encoding
	person := Person{Name: "Jane Doe", Age: 25, IsAdult: true}
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error encoding JSON: ", err)
		return
	}
	fmt.Println("JSON data: ", string(jsonData))

	// decoding
	var newPerson Person
	err = json.Unmarshal(jsonData, &newPerson)
	if err != nil {
		fmt.Println("Error decoding JSON: ", err)
		return
	}

	fmt.Println("Decoded Person: ", newPerson)
}
