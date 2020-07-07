package utils

import "math"

func reverseSlice(a []int) []int {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

// FindShortestPath makes use of dijkstra function
// to return the shortest path from a directed acyclic graph.
// It returns the distance slice and a shortest_path slice.
// g - adjacency list of a weighted graph
// n - the number of nodes in the graph
// s - the index of the starting node ( 0 <= s < n )
// e - the index of the end node (0 <= e < n )
func FindShortestPath(g adjList, n int, s int, e int) ([]int, []int) {
	dist, prev := AStarSearch(g, n, s, e)
	path := make([]int, 0)

	if dist[e] == math.MaxInt64 {
		return path, nil
	}

	// start from the end and loop all the way back to the beginning.
	for pointer := e; pointer != 0; pointer = prev[pointer] {
		path = append(path, pointer)
	}
	shortestPath := reverseSlice(path)

	return dist, shortestPath
}
