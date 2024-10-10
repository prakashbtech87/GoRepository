package main

import "fmt"

// Define the bot interface
type bot interface {
	translate() string
}

// Define the englishTranslator struct
type englishTranslator struct{}

// Define the frenchTranslator struct
type frenchTranslator struct{}

func main() {
	eng := englishTranslator{}
	fre := frenchTranslator{}

	// Call printMessage with different translators
	printMessage(eng)
	printMessage(fre)
}

// printMessage function takes a bot interface and prints the translated message
func printMessage(b bot) {
	fmt.Println(b.translate())
}

// translate method for englishTranslator
func (eng englishTranslator) translate() string {
	return "Hello in English"
}

// translate method for frenchTranslator
func (french frenchTranslator) translate() string {
	return "Bonjour en Français"
}
