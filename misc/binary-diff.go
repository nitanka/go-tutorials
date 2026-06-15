package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func insertNode(node *Node, value int) *Node {
	if node == nil {
		return &Node{value: value}
	}
	if value < node.value {
		fmt.Printf("At node %d: going LEFT for %d\n", node.value, value)
		node.left = insertNode(node.left, value)
	} else {
		fmt.Printf("At node %d: going RIGHT for %d\n", node.value, value)
		node.right = insertNode(node.right, value)
	}
	return node
}

func displayNode(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("At node %d: going LEFT\n", node.value)
	displayNode(node.left)
	fmt.Println("Value is:", node.value)
	fmt.Printf("At node %d: going RIGHT\n", node.value)
	displayNode(node.right)
}
