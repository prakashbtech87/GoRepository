package main

import (
	"fmt"
	"time"
)

func main() {
	//currentTime := time.Now()
	//formattedTime := currentTime.Format("2006-01-02 15:04:05")
	//fmt.Println("Current Time: ", formattedTime)

	// use the "3:04 PM" format for 12 hours clock with AM/PM
	//formattedTime := currentTime.Format("2006-01-02 3:04 PM")
	//fmt.Println("Current Time: ", formattedTime)

	// dataStr := "2024-01-01"
	// parseTime, err := time.Parse("2006-01-02", dataStr)
	// if err == nil {
	// 	fmt.Println("Parsed time: ", parseTime)
	// } else {
	// 	fmt.Println("Error Parsing timr: ", err)
	// }

	currentTime := time.Now()

	// Add 1 day
	newTime := currentTime.Add(24 * time.Hour)

	fmt.Println("Current Time: ", currentTime)
	fmt.Println("New Time: ", newTime)
}
