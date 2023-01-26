package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	duration := time.Since(start)
	fmt.Printf("Elapsed Time: %.5fms\n", duration.Seconds()*1000)
}
