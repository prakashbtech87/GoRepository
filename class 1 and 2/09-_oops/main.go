package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	counter := 0
	for {
		fmt.Println("Infinite loop")
		counter++
		if counter > 5 {
			break
		}

	}

	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Println("\n Index:", index, "Value:", value)
	}

	mesage := "Hello, World!"

	for index, char := range mesage {
		fmt.Printf("\nIndex: %d, Character: %c\n", index, char)
	}

}
