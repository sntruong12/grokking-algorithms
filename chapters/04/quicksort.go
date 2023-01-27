package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	exercise4_1()

	duration := time.Since(start)
	fmt.Printf("Elapsed Time: %.5fms\n", duration.Seconds()*1000)
}

// Write code that will return sum ints in an array
func exercise4_1() {
	numbers := []int{4, 3, 9, 5, 1, 1}

	sum := recursiveSumOfArray(numbers)
	fmt.Println(sum)
}

func recursiveSumOfArray(arr []int) int {
	log.Println(arr)
	if len(arr) == 0 {
		return 0
	} else {
		return arr[0] + recursiveSumOfArray(arr[1:])
	}
}
