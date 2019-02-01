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

// Type is used a lot, here is some more information
// Concrete types: map, struct, int, string.
// Concrete types can be used and meddled with and used to create a value
// Interface type cannot be used to create a value
// Interfaces are not generic types
// Interfaces are implicit - We do not need to declare what uses the interface
// Interfaces are a contract to help us manage types
//  Garbage in -> Garbage out, if our custom types implementation of a function
//  is broken then interfaces wont help us
// Interfaces are tough, Step #1 is understanding how to read them
//  Understand how to read interfaces in the standard lib. Writting your
//  own interfaces is tough and requires experience
