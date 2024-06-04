package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	presentTime := time.Now()

	fmt.Println(presentTime.Format("02-01-2006"))

	createdDate := time.Date(2020, time.April, 23, 3, 2, 0, 0, time.UTC)

	fmt.Println("Crated time:", createdDate.Format("02.01.2006: Monday"))

	numCpu := runtime.NumCPU()

	fmt.Println("Number of cores:", numCpu)

}
