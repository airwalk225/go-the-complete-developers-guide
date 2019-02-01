package main

import "fmt"

func main() {
	// A map is like a struct, but all the keys and values must be the same type
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	// Zero filled map declaration
	var colors2 map[string]string

	// Zero filled map created with make
	colors3 := make(map[string]string)

	// Assignment, you cannot do assignment with name: value, you have to us square braces
	colors["white"] = "#ffffff"
	// You can then assign values to this map, due to the way it was created
	// colors2["white"] = "#ffffff"
	colors3["white"] = "#ffffff"

	// Remove values from the map
	delete(colors, "white")

	printMap(colors)
	printMap(colors2)
	printMap(colors3)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code for %v is %v\n", color, hex)
	}
}

/*
	Why would you use a map over structs
	Maps:
		All keys must be the same type
		All values must be the same type
		Keys are indexed - we can iterate over them
		Use to represent a collection of related properties
		Don't need to know all the keys at compile time
		Reference type!
	Struct
		Values can be a different type
		Keys don't support indexing
		You need to know all the different fields at compile time
		Use ot represent a "thing" with a lot of different properties
		Value type!!
*/
