package main

import "fmt"

func main() {

	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Chiku"
	fruits[3] = "Dragon Fruit"

	fmt.Println("Array Elements:", fruits)

	var vegetables = [4]string{"Tinda", "Carrot", "Corn"}

	fmt.Println("Veggies:", vegetables)
}
