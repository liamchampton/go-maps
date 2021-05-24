package main

import "fmt"

func main() {

	// option 1: map of keys + values of type string
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#00FF00",
		"white": "#ffffff",
	}

	// option 2: map of keys + values of type string
	// good if you want to figure out later on what to add to it
	// var colors map[string]string

	// option 3: map of keys + values of type string
	// colors := make(map[string]string)

	// colors["white"] = "#ffffff"

	// delete element from map
	// delete(colors, "white")

	// fmt.Println(colors)
	// fmt.Println(len(colors))

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
