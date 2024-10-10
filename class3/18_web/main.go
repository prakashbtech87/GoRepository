package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Start of the program")

	// Make a GET request to the given URL

	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	fmt.Println("type of response: %T\n ", response)
	if err != nil {
		fmt.Println("Error making GET request: ", err)
		return
	}
	defer response.Body.Close()

	// read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		return
	}

	// print the response body
	fmt.Println(string(body))
}
