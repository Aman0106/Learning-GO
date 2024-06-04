package main

import "fmt"

func main() {
	fmt.Println("Pointers ahead")

	v := 23

	var p *int = &v

	fmt.Printf("type of p %T\n", p)
	fmt.Println("value of p:", p)
	fmt.Println("value at p:", *p)
}
