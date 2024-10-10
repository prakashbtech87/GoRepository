package main

import "fmt"

func main() {
	var num int = 2
	var ptr *int
	ptr = &num
	fmt.Println("value of variable :", num)
	fmt.Println("value of pointer :", *ptr)

	data := "Go Lang"
	pointer := &data
	fmt.Println("value via pointer :", *pointer)

	var ptr2 *int
	if ptr2 == nil {
		fmt.Println("Nil pointer")
	}

}
