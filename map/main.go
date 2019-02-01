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

	fmt.Println(colors)
	fmt.Println(colors2)
	fmt.Println(colors3)
}
