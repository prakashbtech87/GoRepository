package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println("x is greater then 5")
	} else {
		fmt.Println("x is not greater then 5")
	}

	z := 7
	if z > 10 {
		fmt.Println("z is greater then 10")
	} else if z > 5 {
		fmt.Println("z is greater then 5")
	} else {
		fmt.Println("z is less then 5")
	}

}
