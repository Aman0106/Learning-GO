package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("ENter Value: ")

	// comma ok || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for input:", input)
}
