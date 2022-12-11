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
