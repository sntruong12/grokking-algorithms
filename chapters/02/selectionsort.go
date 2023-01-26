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
		arr = append(arr[:largestIndex], arr[largestIndex+1:]...) // this line basically creates a new array with elements before and after largestIndex - essential omits it from the original
	}

	// arr[:smallestIndex] creates a new slice that contains all elements of arr from the start of the slice to the element before the element at the smallestIndex.
	// arr[smallestIndex+1:] creates a new slice that contains all elements of arr from the element after the element at the smallestIndex to the end of the slice.
	// append(arr[:smallestIndex], arr[smallestIndex+1:]) concatenates these two slices together, effectively omitting the element at the smallestIndex from the resulting slice.
	// ... is the variadic operator, which allows the append function to accept multiple slices as arguments.

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
