package main

import "fmt"

func main() {

	days := []string{"Monday", "Sunday", "Tuesday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// fmt.Println("Using range")

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for _, value := range days {
		fmt.Println(value)
	}
}
