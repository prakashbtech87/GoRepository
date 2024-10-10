package main

import (
	"fmt"
	"time" // Import the time package for working with dates and times
)

func main() {
	// Get the current local time
	currentTime := time.Now()

	// Format the current time as "YYYY-MM-DD HH:MM:SS"
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println("Current Time (24-hour format):", formattedTime)

	// Format the current time as "YYYY-MM-DD HH:MM AM/PM"
	formattedTime = currentTime.Format("2006-01-02 3:04 PM")
	fmt.Println("Current Time (12-hour format):", formattedTime)

	// Example of parsing a date string
	dataStr := "2024-01-01"
	parseTime, err := time.Parse("2006-01-02", dataStr) // Parse a date string
	if err == nil {
		fmt.Println("Parsed time:", parseTime) // Print the parsed time
	} else {
		fmt.Println("Error parsing time:", err) // Print the error if parsing fails
	}

	// Add 1 day (24 hours) to the current time
	newTime := currentTime.Add(24 * time.Hour)

	// Print the current time
	fmt.Println("Current Time:", currentTime)
	// Print the new time after adding 1 day
	fmt.Println("New Time (after adding 1 day):", newTime)
}
