package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Main Function")

	mCh := make(chan int, 2)
	waitGroup := &sync.WaitGroup{}

	waitGroup.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		ch <- 3
		ch <- 6
	}(mCh, waitGroup)

	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("channel:", <-ch)
	}(mCh, waitGroup)

	waitGroup.Wait()
}
