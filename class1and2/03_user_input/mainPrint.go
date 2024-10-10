package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Using bufio.NewReader to read user input from standard input (stdin)
	reader := bufio.NewReader(os.Stdin) // Initialize reader to read from console

	fmt.Println("Enter your name:") // Prompt user for input

	// ReadString reads input until a newline character is encountered
	// The second value (error) is ignored here using the blank identifier '_'
	name, _ := reader.ReadString('\n') // Reads user input including spaces

	// Print the greeting with the entered name
	fmt.Printf("Hello, %s", name) // Outputs the name provided by the user
}
