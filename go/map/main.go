package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4bf745",
		"white": "ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Color", color, "is", "hex", hex)
	}
}
