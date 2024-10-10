package main

import "fmt"

type bot interface {
	translate() string
}
type englishTranslator struct{}
type frenchTranslator struct{}

func main() {
	eng := englishTranslator{}
	fre := frenchTranslator{}
	printMessage(eng)
	printMessage(fre)
}

//	func printMessage(eng englishTranslator){
//	 fmt.Println(eng.translate())
//	}
//
//	func printMessage(fre frenchTranslator){
//	 fmt.Println(fre.translate())
//	}
func printMessage(b bot) {
	fmt.Println(b.translate())
}
func (eng englishTranslator) translate() string {
	return "Hello in english"
}
func (french frenchTranslator) translate() string {
	return "Hello in french"
}
