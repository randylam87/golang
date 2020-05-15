package main

import "fmt"

type colorsMap map[string]string

func main() {
	colors := colorsMap{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c colorsMap) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex+".")
	}
}
