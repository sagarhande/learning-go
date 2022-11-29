package main

import (
	"fmt"
)

func main() {
	// 1st way to declare map

	// colors := map[string]string{
	// 	"red":   "#FF0000",
	// 	"white": "#FFFFFF",
	// }

	// 2nd way to declare map

	//var colors map[string]string  // this will create a nil map i.e. a reference pointing to Nil  DON't USE THIS, USE BELOW

	colors := map[string]string{}

	// 3rd way to declare map

	// colors := make(map[string]string)

	// ADD VALUES TO MAP SAME AS Dictionary in PYTHON

	colors["red"] = "#FF0000"
	colors["white"] = "#FFFFFF"

	fmt.Println(colors)

	// delete key from map

	delete(colors, "red")
	fmt.Println(colors)

}
