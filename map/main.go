package main

import (
	"fmt"
)

func main() {

	// In go, when using a map, since go is a statically typed language, all keys should be of same type and all values should be of the same type.
	// map is a referen e type
	// create a map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	fmt.Println(colors)

	var emptyMap map[int]string
	fmt.Println(emptyMap)

	otherEmptyMap := make(map[string]int)
	fmt.Println(otherEmptyMap)

	// Add elements
	colors["white"] = "#ffffff"
	fmt.Println(colors)

	// Delete map elements
	delete(colors, "red")
	fmt.Println(colors)

	printMap(colors)

}

// Iterate over a map
func printMap(m map[string]string) {
	for color, hex := range m { //key, value
		fmt.Println("Hex code for", color, "is", hex)
	}
}
