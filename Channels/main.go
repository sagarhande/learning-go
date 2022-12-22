package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://amazon.com",
		"https://netflix.com",
		"https://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)

	}

	// for {
	// 	go checkLink(<-c, c)
	// }

	// alternative loop syntax for above

	for l := range c {
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
