package main

import "fmt"

func main() {
    // Variables with different data types
    var name string = "John Doe"  // String variable
    var age int = 25              // Integer variable
    var height float64 = 5.9      // Float variable
    var isMarried bool = true     // Boolean variable

    // Modifying a variable
    isMarried = false             // Changing value of 'isMarried'

    // Print the variables
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Height:", height)
    fmt.Println("Married:", isMarried)

    // Constants in Go
    const pi float64 = 3.14       // Constant value (unchangeable)
    // pi = 6.14                  // Uncommenting this will cause an error as constants can't be reassigned

    fmt.Println("PI:", pi)

    // Shorthand variable initialization
    rollNumber := 101             // Using := for short variable declaration
    fmt.Println("Roll Number:", rollNumber)

    // Visibility of variables in Go based on capitalization
    // If the first letter is uppercase, the variable is exported (public)
    // If the first letter is lowercase, it's private to the package

    // Public variable (accessible outside the package)
    var PublicVariable int = 100

    // Private variable (only accessible within the package)
    var privateVariable int = 200

    // Print public and private variables
    fmt.Println("Public Variable:", PublicVariable)
    fmt.Println("Private Variable:", privateVariable)
}