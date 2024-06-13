package main

import (
	"fmt"
	"net/http"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	// 	go greeter("Hi")
	// 	greeter("bye")

	websites := []string{
		"https://learncodeonline.in",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://aman.dev",
	}

	for _, web := range websites {
		go getStatusCode(web)
		waitGroup.Add(1)
	}

	waitGroup.Wait()
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(1 * time.Second)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer waitGroup.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("%s returned status code %d\n", endpoint, res.StatusCode)
}
