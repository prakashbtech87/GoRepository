package main

import "fmt"

func main() {
	fmt.Println("Start of the program")

	// The function calls with defer will be executed when the surrounding function (main) exits.

	defer fmt.Println("This will be executed first")  // This will be the last deferred function executed
	defer fmt.Println("This will be executed second") // This will be executed second to last

	fmt.Println("middle of the program") // This is executed next

	// The program ends here, and deferred functions will execute in reverse order.
}
