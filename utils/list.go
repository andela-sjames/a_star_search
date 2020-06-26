package utils

type adjList []*LinkedList

type node struct {
	vertex int
	weight int
	next   *node
}

// LinkedList struct defined for representing the graph.
type LinkedList struct {
	name string
	head *node
}

// AddNode adds a node to the linked list data structure.
func (l *LinkedList) AddNode(vertex, weight int) {
	n := &node{
		vertex: vertex,
		weight: weight,
	}
	if l.head == nil {
		l.head = n
	} else {
		currentNode := l.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = n
	}
}

// NewList used to create a new linked list.
func NewList(name string) *LinkedList {
	return &LinkedList{
		name: name,
	}
}
