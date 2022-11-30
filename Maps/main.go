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

	fmt.Println(colors)

	// delete key from map

	delete(colors, "red")
	fmt.Println(colors)

}
