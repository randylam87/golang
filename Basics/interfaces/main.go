package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	// Example: hit a specific api endpoint to get the greeting
	return "Hello"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	// Example: hit a specific api endpoint to get the greeting
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
