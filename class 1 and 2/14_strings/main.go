package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "apple,orange,mango"
	parts := strings.Split(str, ",")

	fmt.Println(parts)

	str2 := "one,two,three,four,five"
	count := strings.Count(str2, "two")
	fmt.Println(count)

	str3 := " Hello Google"
	fmt.Println(str3)
	trimmed := strings.TrimSpace(str3)
	fmt.Println(trimmed)

	str4 := "hello"
	str5 := "world"
	result := strings.Join([]string{str4, str5}, ",")
	fmt.Println(result)

}
