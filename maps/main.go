package main

import "fmt"

func main() {
	colors := make(map[string]string)
	colors["white"] = "#ffffff"
	colors["black"] = "#000000"
	// delete(colors, "white")
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// 	"blue":  "#0000ff",
	// }
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
