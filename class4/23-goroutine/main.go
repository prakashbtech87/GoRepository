package main

import (
	"fmt"
	"time"
)

func display(message string) {
	fmt.Println(message)
}

func main() {
	// Run three different goroutines
	go display("Process 1")
	go display("Process 2")
	go display("Process 3")

	// Sleep main goroutine for 2 seconds
	time.Sleep(2 * time.Second)
}
