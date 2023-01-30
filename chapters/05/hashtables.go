package main

import (
	"fmt"
	"time"
)

// Golang's implementation of hash tables are through Go Maps where you can a map Go type as the Key to another Go type as the Value
// Essentially creating an object with key value pairs
// A Go map type looks like this:
// map[KeyType]ValueType

func main() {
	start := time.Now()

	duration := time.Since(start)
	fmt.Printf("Elapsed Time: %.5fms\n", duration.Seconds()*1000)
}
