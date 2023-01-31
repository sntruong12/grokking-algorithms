package main

import (
	"fmt"
	"time"
)

// breadth-first search allows you to find shortest distance beween two things
// it is algorithm used on graphs

// a shortest-path problem can be solved using the breadth-first search algorithm

// graphs are made up of nodes and edges
// node being a point
// an edge linking nodes to nodes - neighboring nodes as you can call it (nodes that share an edge)
// graphs are a way to model how different things are connected to one another

// this algo helps answer two types questions
// 1. is there a path from node A to node B?
// 2. what's the shortest path from node A to node B?

// queues

func main() {
	start := time.Now()

	duration := time.Since(start)
	fmt.Printf("Elapsed Time: %.5fms\n", duration.Seconds()*1000)
}
