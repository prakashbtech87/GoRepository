package main

import (
	"fmt"     // Importing the fmt package for formatted I/O
	"io"      // Importing the io package for I/O operations
	"os"      // Importing the os package to access OS functions like Stdout
	"strings" // Importing the strings package for string manipulation
)

func main() {
	// Define source using strings.NewReader, which creates a Reader from the given string
	src := strings.NewReader("Hello World")

	// Define destination using os.Stdout, which refers to the standard output (console)
	dst := os.Stdout

	// Call io.Copy to copy data from the source (src) to the destination (dst)
	bytes, err := io.Copy(dst, src)

	// Check if an error occurred during the copy operation
	if err != nil {
		// If there is an error, panic and print the error message
		panic(err)
	}

	// Print a separator line
	fmt.Println("----")

	// Print the number of bytes copied to the destination
	fmt.Printf("the number of bytes copied are %d\n", bytes)
}
