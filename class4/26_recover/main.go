package main

import "fmt"

func mayPanic() {
	panic("a problem occured")
}

func main() {
	fmt.Println("before mayPanic...")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from error", r)
		}
	}()

	mayPanic()

	fmt.Println("after mayPanic...")
}
