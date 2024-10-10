package main

import (
	"fmt"     // Import the fmt package for formatted I/O
	"net/url" // Import the net/url package for URL parsing and manipulation
)

func main() {
	// Parse a URL string into a URL object
	myURL := "https://example.com:8080/path/to/resource?name=Jane+Doe&age=25"
	parsedURL, err := url.Parse(myURL) // Parse the URL
	if err != nil {
		fmt.Println("Error parsing URL:", err) // Handle parsing error
		return
	}

	// Accessing URL components
	fmt.Println("Scheme:", parsedURL.Scheme)     // Print the scheme (protocol)
	fmt.Println("Host:", parsedURL.Host)         // Print the host (domain and port)
	fmt.Println("Path:", parsedURL.Path)         // Print the path
	fmt.Println("RawQuery:", parsedURL.RawQuery) // Print the raw query string

	// Modify the URL components
	parsedURL.Scheme = "http"         // Change the scheme to HTTP
	parsedURL.Host = "newhost.com"    // Change the host to newhost.com
	parsedURL.Path = "/newpath"       // Change the path to /newpath
	parsedURL.RawQuery = "newquery=1" // Change the query string

	// Construct the modified URL
	newURL := parsedURL.String()         // Convert the modified URL object back to a string
	fmt.Println("Modified URL:", newURL) // Print the modified URL
}
