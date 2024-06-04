package main

import "fmt"

func main() {
	aman := User{"Aman", "asingh@gmail.com", 9121}

	fmt.Println("My User:", aman)
	fmt.Printf("My User:%+v", aman)
	fmt.Printf("\nName %v Email %v", aman.Name, aman.Email)
}

type User struct {
	Name   string
	Email  string
	Number int
}
