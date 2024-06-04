package main

import "fmt"

func main() {
	count := 11

	result := "All good"

	if count > 10 {
		result = "not good"
	}

	fmt.Println(result)
}
