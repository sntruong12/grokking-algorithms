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

// base case for quick sort is empty array or array with 1 element
// arrays with 2 elements easy to quick sort
// check if first element is smaller than the second
// if it isn't then swap them

// quick sort inner workings
// select an element from array - this is called the pivot
// now find all elements smaller than pivot
// then find all elements larger than pivot
// this is called partitioning
// you have sub array with all smaller values than pivot
// and the same for larger values than pivot
// note at this point the sub arrays are not sorted

// how do you sort the sub arrays?
// call quicksort recursively on the two sub arrays
// append left + pivot + right to have sorted array

// ChaptGPT response
// QuickSort is a divide and conquer sorting algorithm that uses a pivot element to partition the array into smaller sub-arrays, and then recursively sorts each sub-array. The pivot element is chosen as the last element of the array in this example.

// The basic steps of the QuickSort algorithm are as follows:

// Select an element in the array as the pivot element.
// Partition the array into two sub-arrays, one with elements less than or equal to the pivot and one with elements greater than the pivot.
// Recursively sort the sub-array of elements less than or equal to the pivot.
// Recursively sort the sub-array of elements greater than the pivot.
// Merge the sorted sub-arrays and the pivot to form the final sorted array.

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)-1]
	left := make([]int, 0, len(arr))
	right := make([]int, 0, len(arr))

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] <= pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)
}
