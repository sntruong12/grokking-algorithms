package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	// exercise4_1()
	// exercise4_2()
	exercise4_3()

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

func exercise4_3() {
	max := betterFindMaxNumber([]int{123, 23, 1231, 455, 5557, 863, 23, 5557, 551, 60, 790}, 0)
	fmt.Println(max)
	// max = findMaxNumber([]int{})
	// fmt.Println(max)
	// max = findMaxNumber([]int{123})
	// fmt.Println(max)
	// max = findMaxNumber([]int{12, 3})
	// fmt.Println(max)
}

func findMaxNumber(arr []int) int {
	if len(arr) == 0 {
		return 0
	} else if len(arr) == 1 {
		return arr[0]
	} else if arr[0] > arr[1] {
		return findMaxNumber(append(arr[:1], arr[2:]...)) // remove the 1 index in arry
	} else {
		return findMaxNumber(arr[1:]) // remove the 0 index in array
	}
}

// better way of writing it
func betterFindMaxNumber(arr []int, max int) int {
	if len(arr) == 0 {
		return max
	}
	if arr[0] > max {
		max = arr[0]
	}
	return betterFindMaxNumber(arr[1:], max)
}

// base case of binary search is if list is zero or one, then we know not to continue a recursive function as we know the answer
// recursive case is based on what we are searching for, if it's greater than or less than the search query then we will continue the search function on the truthy case meaning first half or second half.
