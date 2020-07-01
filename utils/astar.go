package utils

/**

**/
func aStarSearch(g adjList, n int, s int, e int) ([]int, []int) {
	// g - adjacency list of a weighted graph
	// n - the number of nodes in the graph
	// s - the index of the starting node ( 0 <= s < n )
	// e - the index of the end node ( 0 <= e < n )

	// visited := make([]bool, n)
	distance := make([]int, n)

	// keep track of the previous node we took
	// to get to the current node
	previous := make([]int, n)

	return distance, previous
}
