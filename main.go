package main

import (
	"fmt"

	utils "github.com/andela-sjames/a_star_search/utils"
)

// A * search
func main() {
	nodezero := utils.NewList("nodezero")
	nodezero.AddNode(1, 4)
	nodezero.AddNode(2, 1)

	nodeone := utils.NewList("nodeone")
	nodeone.AddNode(3, 1)

	nodetwo := utils.NewList("nodetwo")
	nodetwo.AddNode(1, 2)
	nodetwo.AddNode(3, 5)

	nodethree := utils.NewList("nodethree")
	nodethree.AddNode(4, 3)

	nodefour := utils.NewList("nodefour")

	fmt.Println(nodefour)
}
