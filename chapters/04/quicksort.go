package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	// exercise4_1()
	exercise4_2()

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

// Write a recursive funciton to count the number of items in a list
func exercise4_2() {
	count := recursiveCountItemsInList([]int{1, 23, 43, 455, 513, 862, 7544})
	fmt.Println(count)
}

func recursiveCountItemsInList(arr []int) int {
	if len(arr) == 0 {
		return 0
	} else {
		return 1 + recursiveCountItemsInList(arr[1:])
	}
}
