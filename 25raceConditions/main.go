package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Main function")

	score := []int{0}

	var waitGroup = &sync.WaitGroup{}
	mut := &sync.Mutex{}

	waitGroup.Add(3)

	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		fmt.Println("1 Go")
		mu.Lock()
		score = append(score, 1)
		mut.Unlock()
		defer wg.Done()
	}(waitGroup, mut)

	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		fmt.Println("2 Go")
		mu.Lock()
		score = append(score, 2)
		mut.Unlock()
		defer wg.Done()
	}(waitGroup, mut)

	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
		fmt.Println("3 Go")
		mu.Lock()
		score = append(score, 3)
		mut.Unlock()
		defer wg.Done()
	}(waitGroup, mut)

	waitGroup.Wait()

	fmt.Println("SLice:", score)
}
