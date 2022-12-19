package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	/*
		// Asignment 1
		s := squere{side: 10}
		t := triangle{base: 7, height: 9}

		printArea(s)
		printArea(t)
	*/

	/* Assignment 2*/
	if len(os.Args) != 2 {
		fmt.Println("Wrong number of command line arguments received")
		os.Exit(1)
	}
	fileName := os.Args[1]

	f, err := os.Open(fileName)

	if err != nil {
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)

}
