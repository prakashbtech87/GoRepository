package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Parse a URL string into a URL object
	myURL := "https://example.com:8080/path/to/resource?name=Jane+Doe&age=25"
	parsedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Error parsing URL: ", err)
		return
	}

	// Accessing URL components
	fmt.Println("Scheme: ", parsedURL.Scheme)
	fmt.Println("Host: ", parsedURL.Host)
	fmt.Println("Path: ", parsedURL.Path)
	fmt.Println("RawQuery: ", parsedURL.RawQuery)

	// modify the URL components
	parsedURL.Scheme = "http"
	parsedURL.Host = "newhost.com"
	parsedURL.Path = "/newpath"
	parsedURL.RawQuery = "newquery=1"

	// construct the modified URL
	newURL := parsedURL.String()
	fmt.Println("Modified URL: ", newURL)

}
