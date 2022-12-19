package main

import "fmt"

type bot interface {
	getGreeting() string
}

type hindiBot struct{}
type spanishBot struct{}

func main() {

	eb := hindiBot{}
	sp := spanishBot{}

	printGreetings(eb)
	printGreetings(sp)

	/*
		// Asignment 1
		// sagar := person{firstName: "Sagar", lastName: "Hande"}
		s := squere{side: 10}
		t := triangle{base: 7, height: 9}

		printArea(s)
		printArea(t)
	*/

}

func printGreetings(b bot) {
	fmt.Println(b.getGreeting())
}

func (h hindiBot) getGreeting() string {
	return "Namaskar!"
}

func (s spanishBot) getGreeting() string {
	return "Hola!"
}
