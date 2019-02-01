package main

import "fmt"

type bot interface {
	getGreeting() string
}

// This is a test interface to show a more advanced use
type botAdvanced interface {
	// This interface contract requires arguments of string and int and will return string and error
	getAdvancedGreeting(string, int) (string, error)
	// This contract will return a type of float 64
	getBotVersion() float64
	// This contact will accept an argument of user and return a type of string
	respondToUser(user) string
}

type user struct{}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
