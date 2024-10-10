package main

import "fmt"

func main() {

	// declare and initialize a slice
	numbers := []int{1, 2, 3, 4, 5}

	// // access and print slice elements
	fmt.Println("Slice elements: ", numbers)

	// // length of the slice
	fmt.Println("Length of the slice: ", len(numbers))

	// creating a slice without specifying the capacity

	nums := []int{1, 2, 3, 4}
	fmt.Println("Slice elements:", nums)
	fmt.Println("Length elements:", len(nums))
	fmt.Println("Capacity elements:", cap(nums))

	numsCheck := make([]int, 3, 5)
	numsCheck = append(numsCheck, 4)
	numsCheck = append(numsCheck, 5)
	numsCheck = append(numsCheck, 6)

	numsCheck = append(numsCheck, 4)
	numsCheck = append(numsCheck, 5)
	numsCheck = append(numsCheck, 6)
	numsCheck = append(numsCheck, 4)
	numsCheck = append(numsCheck, 5)

	// print the sliced array
	fmt.Println("\n\nSliced array: ", numsCheck)
	fmt.Println("Length : ", len(numsCheck), "Capacity : ", cap(numsCheck))

}
