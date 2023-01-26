package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	unsortedInts := []int{99, 354, 1, 52, 9000, 621, 29, 505, 900, 505}
	sortedInts := selectionSort(unsortedInts)

	fmt.Println(sortedInts)

	duration := time.Since(start)
	fmt.Printf("Elapsed Time: %.5fms\n", duration.Seconds()*1000)
}

func selectionSort(arr []int) []int {
	newArr := make([]int, 0)

	for range arr {
		// smallestIndex := findSmallest(arr)
		// newArr = append(newArr, arr[smallestIndex])
		// arr = append(arr[:smallestIndex], arr[smallestIndex+1:]...)
		largestIndex := findLargest(arr)
		newArr = append(newArr, arr[largestIndex])
		arr = append(arr[:largestIndex], arr[largestIndex+1:]...)
	}

	return newArr
}

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0

	for i, v := range arr {
		if v < smallest {
			smallest = v
			smallestIndex = i
		}
	}

	return smallestIndex
}

func findLargest(arr []int) int {
	largest := arr[0]
	largestIndex := 0

	for i, v := range arr {
		if v > largest {
			largest = v
			largestIndex = i
		}
	}

	return largestIndex
}
