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

// queues are lists of things to do in a sequential order
// you can't access random elements in the queue
// there are only two operations enqueue (add an item to the queue) and dequeue (take an item off the queue)
// the queue is called a FIFO data structure - first in first out
// in contrast a stack is a LIFO data structure - last in first out

func main() {
	start := time.Now()

	duration := time.Since(start)
	fmt.Printf("Elapsed Time: %.5fms\n", duration.Seconds()*1000)
}

type apartmentResidents map[string]person

type person struct {
	name string
	causedAFlood bool
	neighbors []person
}

var (
	centerline := apapartmentResidents{
		"Yen": {
			name: "Yen"
			causedAFlood: false,
			neighbors: []{
				"4UnitsApartmentFlooder": {
					name: "unknown",
					causedAFlood: true,
				}
				"Terrance":{
					causedAFlood: false
				}
			}
		},
		"4UnitsApartmentFlooder": {
			causedAFlood: true,
		},
		"Terrance":{
			causedAFlood: false
		},
	}
)

func breadthFirstSearch() {

}