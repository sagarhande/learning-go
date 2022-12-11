package main

import (
	"fmt"
)

func main() {
	// 1st way to declare map

	colors := map[string]string{
		"red":   "#FF0000",
		"white": "#FFFFFF",
		"green": "#00FF00",
	}

	// delete key from map
	// delete(colors, "red")
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hexCode := range c {
		fmt.Println("Hex code for color ", color, " is ", hexCode)
	}
}
