package main

import "fmt"

func main() {
	aman := User{"Aman", "asingh@gmail.com", 9121}
	pankaj := User{"Pankaj", "asingh@gmail.com", 9121}

	aman.printName()
	pankaj.printName()
}

type User struct {
	Name   string
	Email  string
	Number int
}

// It is using pass by value a new struct is created

func (u User) printName() {
	fmt.Println("User name is:", u.Name)
}
