package main

import "fmt"

func main() {
	fmt.Println("Start of the program")

	// the fuction inside defer will be executed when surrounding function exits

	//defer fmt.Println("This will be printed at the end of the program")

	defer fmt.Println("This will be executed second")

	defer fmt.Println("This will be executed first")

	fmt.Println("middle of the program")
}
