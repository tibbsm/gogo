package main

import "fmt"

func main() {
	// keys are [string] and values are string
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#ffff00",
	// }

	// var colors map[string]string

	colors := make(map[string]string)

	colors["white"] = "#ffffff"
	colors["black"] = "#000000"
	colors["blue"] = "#??????"

	// delete(colors, "white")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, " - ", hex)
	}
}
