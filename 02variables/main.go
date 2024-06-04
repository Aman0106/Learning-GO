package main

import "fmt"

func main() {
	var name string = "Aman"
	fmt.Println(name)
	fmt.Printf("Variable of type %T \n", name)

	var smallInt uint8 = 244
	fmt.Println(smallInt)
	fmt.Printf("Variable of type %T \n", smallInt)

	var intAlias int = 2442323223232322221
	fmt.Println(intAlias)
	fmt.Printf("Variable of type %T \n", intAlias)
}
