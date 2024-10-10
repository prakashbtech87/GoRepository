// Go program to illustrate how
// to create an anonymous Goroutine
// package main

// import (
// 	"fmt"
// 	"time"
// )

// // Main function
// func main() {

// 	fmt.Println("Welcome!! to Main function")

// 	// Creating Anonymous Goroutine
// 	go func() {
// 		fmt.Println("Welcome!! to First")
// 	}()

// 	time.Sleep(5 * time.Second)
// 	fmt.Println("GoodBye!! to Main function")
// }

package main

import (
	"fmt"
	"time"
)

func display(message string) {
	fmt.Println(message)
}

func main() {
	// run two differet goroutine
	go display("Process 1")
	go display("Process 2")
	go display("Process 3")

	// tp sleep main goroutine for 2 sec
	time.Sleep(2 * time.Second)
}
