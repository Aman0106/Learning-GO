package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(addNumbers(numbers[:3]...))

}

func addNumbers(nums ...int) int {
	total := 0

	for _, val := range nums {
		total += val
	}

	return total
}
