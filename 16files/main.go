package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	content := "This is saved to the file in  using go"

	file, err := os.Create("./myLogFile.txt")

	if err != nil {
		panic(err)
		// fmt.Println("Error occured as:", err)
	}

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is: ", length)

	defer file.Close()

	readFile("file.Name()")
}

func readFile(filename string) {
	content, err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("contents are:", string(content))

}

func checkNilErr(err error) {
	if err != nil {
		fmt.Println("Error is:", err)
	}
}
