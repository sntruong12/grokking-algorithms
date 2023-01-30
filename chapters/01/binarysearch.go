package main

import (
	"fmt"
	"time"
)

// binary search is an algorithm where input is a sorted list of elements
// returns position of what you are searching for
// or returns null

// eliminates half the number every time with this algo

// which essentially translate to log2 N
// list of 8 numbers would take at most 3 tries because log2 8 = 3
// 2 raised to the what number is 8? 3
// list of 1024 would take at most 10 tries because log2 1024 = 10
// 2 raised to the what number is 1024? 10

func main() {
	start := time.Now()

	item := 159354
	list := createSortedList(205001)
	// position, err := binarySearch(list, item)
	position, err := simpleSearch(list, item)
	if err != nil {
		fmt.Printf("%d not found in list\n", item)
	} else {
		fmt.Println("position is: ")
		fmt.Println(position)
	}

	duration := time.Since(start)
	fmt.Printf("Elapsed Time: %.5fms\n", duration.Seconds()*1000)
}

func createSortedList(size int) []int {
	list := make([]int, 0)
	for i := 0; i < size; i++ {
		list = append(list, i)
	}

	return list
}

func binarySearch(sortedList []int, searchFor int) (int, error) {
	// low and high keeps track of which part of the sortedList you'll search in
	low := 0
	high := len(sortedList) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := sortedList[mid]
		if guess == searchFor {
			return mid, nil
		}

		if guess > searchFor {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0, fmt.Errorf("%d not found in list", searchFor)
}

func simpleSearch(sortedList []int, searchFor int) (int, error) {
	for i, v := range sortedList {
		if v == searchFor {
			return i, nil
		}
	}

	return 0, fmt.Errorf("%d not found in list", searchFor)
}
