package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// Example 1
	// create a new file or truncate an existing one
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}
	defer file.Close()

	// initial content to be added to the file
	initialContent := "Hello, this is the initial content of the file"

	// write the initial content to the file using io.WriteString
	_, err = io.WriteString(file, initialContent)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}

	fmt.Println("File created and written to successfully")

	// Example 2
	// open a file
	file, err = os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	// create a buffer to read the file content
	buffer := make([]byte, 1024)

	// read the file content
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break // end of file reached
		}
		if err != nil {
			fmt.Println("Error reading file: ", err)
			return
		}

		// process the read content
		fmt.Print(string(buffer[:n]))
	}

	// Example 3
	// read the entire file into byte slice
	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	// process the file content
	fmt.Println(string(content))
}
