package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#ffgreen",
		"white": "#ffwhite",
	}

	//var colors map[string]string

	//colors := make(map[string]string)

	//colors["white"] = "#ffwhite"
	//delete(colors, "white")

	printMap(colors)
}

func printMap(colors map[string]string) {
	for color, hex := range colors {
		fmt.Println("Hex Code for", color, "is", hex)
	}
}
