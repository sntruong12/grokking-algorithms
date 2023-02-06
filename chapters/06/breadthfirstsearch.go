package main

import (
	"fmt"
	"strings"
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

	found := breadthFirstSearch("yen", "simpson")
	fmt.Println(found)

	duration := time.Since(start)
	fmt.Printf("Elapsed Time: %.5fms\n", duration.Seconds()*1000)
}

func breadthFirstSearch(startName, name string) bool {
	searchQueue := make([]string, 0)

	graph := map[string][]string{
		"yen":        {"flooder", "josie", "son", "terrance"},
		"son":        {"josie", "terrance"},
		"terrance":   {"yen"},
		"flooder":    {"yen", "jimbomango"},
		"josie":      {"bob"},
		"bob":        {"josie"},
		"sally":      {},
		"tishawn":    {},
		"jimbomango": {"taylor", "simpson"},
	}

	searchQueue = graph[startName]

	searched := map[string]bool{}

	for len(searchQueue) > 0 {
		person := searchQueue[0]
		searchQueue = searchQueue[1:]

		if _, ok := searched[person]; !ok {
			if strings.EqualFold(person, name) {
				fmt.Println("found this mango seller", person)
				return true
			}

			searchQueue = append(searchQueue, graph[person]...)
			searched[person] = true
		}
	}

	fmt.Println("unable to the mango seller", name)
	return false
}

// contains checks if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
