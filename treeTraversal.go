/*
the Tree coming in will be an array
*/

package main

import "fmt"

type node struct {
	name     string
	contents string
	left     *node
	right    *node
}

func traverse(root *node) error {
	// traverse to the left
	printNode(root)
	left := root.left
	if err := traverse(left); err != nil {
		return err
	}
	//traverse to the right
	right := root.right
	if err := traverse(right); err != nil {
		return err
	}
	return nil
}

func (n *node) printNode() {
	fmt.Printf("node %v has value %v \n", n, n.contents)
}

func main() {
	node1 := node{
		name:     "node1",
		contents: "contents of node1",
		left:     node2,
		right:    node3,
	}
	node2 := node{
		name:     "node2",
		contents: "contents of node2",
		left:     nil,
		right:    nil,
	}
	node3 := node{
		name:     "node3",
		contents: "contents of node3",
		left:     nil,
		right:    nil,
	}

	traverse(node1)
}
