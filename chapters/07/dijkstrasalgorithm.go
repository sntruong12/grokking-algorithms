package main

import (
	"fmt"
	"math"
	"time"
)

// in previous chapter we found the shortest path in a graph (nodes connected by edges)
// what about finding the fastest or smallest total weighted path? - use the Dijkstra's algorithm

// 4 main steps in Dijkstra's algorithm
// ---
// 1. find the cheapest node - node you can get to in least amount of time
// 2. update the costs of the neighbors of this node
// 3. repeat until you've done this for every node in graph
// 4. calculate the final path

// in this algorithm, you assign a number or weight to each segment (edge) then find the path with the smallest weight
// each edge in the graph has a number associated to it - aka weights!
// a graph with weights is called a weighted graph
// breadth-first search is used on unweighted graphs
// conversely, to find shortest path on weighted graphs use this Dijkstra's

// graphs can also have cycles - you can start at a node and end up back at the same node - that's a cycle
// Dijkstra's algorithm only works with directed acyclic graphs or graphs that don't contain cycles

// create a table to track 3 variables, parent, node, cost
// can't use this algorithm if negative weight edges (cost)
// learn bellman-ford algorithm for graphs that contain negative weight edges

const infinity = math.MaxInt32

func main() {
	start := time.Now()

	dijkstras()

	duration := time.Since(start)
	fmt.Printf("Elapsed Time: %.5fms\n", duration.Seconds()*1000)
}

func dijkstras() {
	graph := map[string]map[string]int{}
	graph["start"] = map[string]int{}
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2

	graph["a"] = map[string]int{}
	graph["a"]["fin"] = 1

	graph["b"] = map[string]int{}
	graph["b"]["a"] = 3
	graph["b"]["fin"] = 5

	graph["fin"] = map[string]int{}

	fmt.Println(graph)

	costs := map[string]int{}
	costs["a"] = 6
	costs["b"] = 2
	costs["fin"] = infinity

	fmt.Println("costs")
	fmt.Println(costs)

	parents := map[string]string{}
	parents["a"] = "start"
	parents["b"] = "start"
	parents["fin"] = ""

	fmt.Println("parents")
	fmt.Println(parents)

	processed := map[string]bool{}

	node := findLowestUnprocessedCostNode(costs, processed) // find lowest code nost that you haven't processed yet
	for node != "" {                                        // if processed all the nodes - loop will exit
		cost := costs[node]
		neighbors := graph[node]

		for k := range neighbors { // go thru all neighbors of this node
			newCost := cost + neighbors[k]
			if costs[k] > newCost { // if it's cheaper to get to this neighbor by going thru this node ...
				costs[k] = newCost // ... update the cost for this node
				parents[k] = node  // this node becomes the new parent for this neighbor
			}
		}
		processed[node] = true                                 // mark node as processed
		node = findLowestUnprocessedCostNode(costs, processed) // find the next node to process and loop
	}

	fmt.Println("finished costs and parents")
	fmt.Println("==================================")
	fmt.Println(costs)
	fmt.Println(parents)
}

func findLowestUnprocessedCostNode(costs map[string]int, processed map[string]bool) string {
	lowest := infinity
	node := ""
	for k, v := range costs { // go thru each node
		if _, ok := processed[k]; !ok && v < lowest { // if it's the lowest cost so far and hasn't been processed yet ..
			lowest = v // ... set it as the new lowest cost node
			node = k
		}
	}

	return node
}
