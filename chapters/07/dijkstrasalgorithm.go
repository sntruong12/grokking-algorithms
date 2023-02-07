package main

import (
	"fmt"
	"time"
)

// in previous chapter we found the shortest path in a graph (nodes connected by edges)
// what about finding the fastest or smallest total weighted path? - use the Dijkstra's algorithm

// 4 main steps in Dijkstra's algorithm
// ---
// 1.  find the cheapest node - node you can get to in least amount of time
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

func main() {
	start := time.Now()

	duration := time.Since(start)
	fmt.Printf("Elapsed Time: %.5fms\n", duration.Seconds()*1000)
}
