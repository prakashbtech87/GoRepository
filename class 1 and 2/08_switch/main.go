package main

import "fmt"

func main() {

	// Example 1
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid day")
	}

	// Example 2
	month := "June"
	switch month {
	case "January", "February", "March":
		fmt.Println("Winter")
	case "April", "May", "June":
		fmt.Println("Spring")
	default:
		fmt.Println("other seasons")
	}

	// Example 3 - swicth with expression
	temperatur := 25
	switch {
	case temperatur < 0:
		fmt.Println("Freezing")
	case temperatur >= 0 && temperatur < 10:
		fmt.Println("Very cold")
	case temperatur >= 10 && temperatur < 20:
		fmt.Println("Cold")
	case temperatur >= 20 && temperatur < 30:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}

}
