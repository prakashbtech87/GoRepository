package main

import (
	"fmt"
	"strings" // Import the strings package for string manipulation functions
)

func main() {
	// Example 1: Splitting a string into parts
	str := "apple,orange,mango"
	parts := strings.Split(str, ",") // Split the string by the delimiter ","
	fmt.Println(parts)               // Print the resulting slice of parts

	// Example 2: Counting occurrences of a substring
	str2 := "one,two,three,four,five"
	count := strings.Count(str2, "two") // Count how many times "two" appears in the string
	fmt.Println(count)                  // Print the count (should print 1)

	// Example 3: Trimming whitespace from a string
	str3 := " Hello Google"
	fmt.Println(str3)                  // Print the original string with leading space
	trimmed := strings.TrimSpace(str3) // Remove leading and trailing whitespace
	fmt.Println(trimmed)               // Print the trimmed string

	// Example 4: Joining multiple strings
	str4 := "hello"
	str5 := "world"
	result := strings.Join([]string{str4, str5}, ",") // Join str4 and str5 with a comma
	fmt.Println(result)                               // Print the resulting string
}
