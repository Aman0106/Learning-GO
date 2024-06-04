package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com"

func main() {

	fmt.Println("Web Request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// fmt.Printf("Response type: %T\n", response)
	// fmt.Printf("Response : %v", response)

	defer response.Body.Close()

	output, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Body :\n", string(output))
}
