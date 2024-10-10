package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// define scource
	src := strings.NewReader("Hello World")

	// defining destination using Stdout
	dst := os.Stdout

	// Calling Copy method with its parameter
	bytes, err := io.Copy(dst, src)

	// if error is not null
	if err != nil {
		//fmt.Println(err)
		panic(err)
	}
	fmt.Println("----")
	// print output
	fmt.Printf("the number of bytes copied are %d\n", bytes)
}
