package main

import "fmt"

// function without parameter and return type
func simpleFunction() {
	fmt.Println("Simple Function")
}

// function with parameter and return type
func add(a int, b int) int {
	// c:=a + b
	// return c
	return a + b
}
func add2(a, b, c, d, e, f, g int) int {
	return a + b + c + d + e + f + g
}

// function with name return variable
func multiply(a, b int) (result int) {
	result = a * b
	return
}
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
func main() {
	simpleFunction()
	sum := add(10, 20)
	fmt.Println("Sum:", sum)
	sum2 := add2(10, 20, 30, 40, 50, 60, 70)
	fmt.Println("Sum2:", sum2)
	multiplyResult := multiply(10, 20)
	fmt.Println("Multiply:", multiplyResult)
	divideResult, _ := divide(10, 5)
	fmt.Println("Divide:", divideResult)
	//divideResult, err := divide(10, 5)
	//fmt.Println("Divide:", divideResult)
}
