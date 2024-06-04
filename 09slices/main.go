package main

import "fmt"

func main() {

	var fruitist = []string{"Apple", "orange"}

	fmt.Println("Elements:", fruitist)
	fmt.Printf("Type:%T\n", fruitist)
	fmt.Println("Lenth:", len(fruitist))

	fruitist = append(fruitist, "Mango", "Banana")

	fmt.Println("After altration:")

	fmt.Println("Elements:", fruitist)
	fmt.Printf("Type:%T\n", fruitist)
	fmt.Println("Lenth:", len(fruitist))

	fruitist = append(fruitist[1:])

	fmt.Println("\nAfter Slicing")

	fmt.Println("Elements:", fruitist)
	fmt.Printf("Type:%T\n", fruitist)
	fmt.Println("Lenth:", len(fruitist))

	highScore := make([]int, 4)

	highScore[0] = 99
	highScore[1] = 9
	highScore[2] = 79
	highScore[3] = 67

	fmt.Println("\nNew Highscores")

	fmt.Println("Elements:", highScore)
	fmt.Printf("Type:%T\n", highScore)
	fmt.Println("Lenth:", len(highScore))

}
