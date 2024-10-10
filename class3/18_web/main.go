package main

import (
	"fmt"       // Import the fmt package for formatted I/O
	"io/ioutil" // Import the ioutil package for reading I/O
	"net/http"  // Import the net/http package for making HTTP requests
)

func main() {
	fmt.Println("Start of the program")

	// Make a GET request to the specified URL
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	// Print the type of response received
	fmt.Printf("Type of response: %T\n", response)

	// Handle error if the request fails
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}

	defer response.Body.Close() // Ensure the response body is closed when done

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)

	// Handle error if reading the response body fails
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response body as a string
	fmt.Println(string(body))
}
