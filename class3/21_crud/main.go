package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   // no `json tag specified`
}

func PerformGetRequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"
	response, err := http.Get(myUrl)
	if err != nil {
		fmt.Println("Error making GET request: ", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: failed to fetch data:", response.StatusCode)
		return
	}

	var todo Todo
	err = json.NewDecoder(response.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding JSON: ", err)
		return
	}

	fmt.Println("Todo Id: ", todo.ID)
	fmt.Println("Todo Title: ", todo.Title)
	fmt.Println("Todo Completed: ", todo.Completed)
}

func PerformPostRequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/todos"

	todo := Todo{
		UserId:    1,
		ID:        201,
		Title:     "Learn Go!",
		Completed: false,
	}

	// convcert the todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error encoding JSON: ", err)
		return
	}

	// convert the SJON byte slice to a string
	jsonStr := string(jsonData)

	// create an io.Reader from the string
	jsonReader := strings.NewReader(jsonStr)

	// send the POST request
	resp, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error making POST request: ", err)
		return
	}
	defer resp.Body.Close()

	// Print the response status code
	fmt.Println("Response Status: ", resp.Status)
}

func PerformPutRequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/todos"

	todo := Todo{
		UserId:    1,
		ID:        1,
		Title:     "Update Todo!",
		Completed: true,
	}

	// convcert the todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error encoding JSON: ", err)
		return
	}

	// create a PUT request
	req, err := http.NewRequest(http.MethodPut, myUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating PUT request: ", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// send the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making PUT request", err)
	}

	defer resp.Body.Close()
	fmt.Println("Response Status: ", resp.Status)
}

func PerformDeleteRequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

	// create a DELETE request
	req, err := http.NewRequest(http.MethodDelete, myUrl, nil)
	if err != nil {
		fmt.Println("Error creating DELETE request: ", err)
		return
	}

	// send the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making DELETE request: ", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response Status: ", resp.Status)
}

func main() {
	//PerformGetRequest()
	//PerformPostRequest()
	//PerformPutRequest()
	PerformDeleteRequest()
}
