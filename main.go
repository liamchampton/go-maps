package main

import "fmt"

func main() {

	// map of keys + values of type string
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#00FF00",
	}

	fmt.Println(colors)
	fmt.Println(len(colors))
}
