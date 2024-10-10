package main

import "fmt"

// func main() {
// 	var myarr [3]string

// 	myarr[0] = "first"
// 	myarr[1] = "second"
// 	myarr[2] = "third"

// 	fmt.Println("Elements of Array:")
// 	fmt.Println("Element 1:", myarr[0])

// 	fmt.Println("Elements of Array:")
// 	fmt.Println("Element 2:", myarr[5])
// }

func entry(lang *string, aname *string) {
	if lang == nil {
		panic("Error: language cannot be nil")
	}

	if aname == nil {
		panic("Error: author name cannot be nil")
	}

	fmt.Printf("Author Language: %s\n Author Name: %s\n", *lang, *aname)
}

func main() {
	A_Lang := "GO Lang"

	entry(&A_Lang, nil)
}
