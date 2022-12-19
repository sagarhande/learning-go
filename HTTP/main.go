package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	res, err := http.Get("https://www.google.com/")

	if err != nil {
		fmt.Println("Error ocurred: ", err)
		os.Exit(1)
	}

	// fmt.Println(res)

	io.Copy(os.Stdout, res.Body)

}
