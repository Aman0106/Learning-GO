package main

import "fmt"

func main() {

	languages := make(map[string]string)

	languages["js"] = "javaScript"
	languages["kt"] = "Kotlin"
	languages["go"] = "GO"

	fmt.Println("All Data:", languages)
	fmt.Println("kt Data:", languages["kt"])

	delete(languages, "js")
	fmt.Println("All Data:", languages)

	for _, value := range languages {
		fmt.Println("Value:", value)
	}
}
