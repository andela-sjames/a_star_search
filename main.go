package main

import (
	"fmt"

	utils "github.com/andela-sjames/a_star_search/utils"
)

// A * search
func main() {
	nodezero := utils.NewList("nodezero", 8)
	nodezero.AddNode(1, 4)
	nodezero.AddNode(2, 1)

	nodeone := utils.NewList("nodeone", 4)
	nodeone.AddNode(3, 1)

	nodetwo := utils.NewList("nodetwo", 7)
	nodetwo.AddNode(1, 2)
	nodetwo.AddNode(3, 5)

	nodethree := utils.NewList("nodethree", 2)
	nodethree.AddNode(4, 3)

	nodefour := utils.NewList("nodefour", 0)

	fmt.Println(nodefour)
	adjlist := []*utils.LinkedList{nodezero, nodeone, nodetwo, nodethree, nodefour}

	n := len(adjlist)
	s := 0
	e := 4

	dist, path := utils.AStarSearch(adjlist, n, s, e)

}
