package main

import (
	"fmt"
	"strconv"
)

func main() {
	// numeric type conversion
	var integerNum int = 40
	var floatNum float64 = float64(integerNum)
	fmt.Println("Float number:", floatNum)
	// string conversion
	var number int = 50
	str := strconv.Itoa(number) // integer to string
	fmt.Println("String:", str)
	strNum := "123"
	num, err := strconv.Atoi(strNum) // string to integer
	fmt.Println("Number:", num)
	fmt.Println("Error:", err)
	// convert string to float
	str2 := "3.14"
	num2, err2 := strconv.ParseFloat(str2, 64)
	if err == nil {
		fmt.Println("Float number:", num2)
	} else {
		fmt.Println(err2)
	}
}
