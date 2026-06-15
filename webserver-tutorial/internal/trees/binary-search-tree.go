package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BST struct {
	node *Node
}

func newNodeCreate(element int) *Node {
	temp := &Node{value: element}
	return temp
}

// Function to insert a node with a specific value in the tree
func insertNode(node *Node, value int) *Node {
	if node == nil {
		return newNodeCreate(value)
	}
	if value < node.value {
		node.left = insertNode(node.left, value)
	} else if value > node.value {
		node.right = insertNode(node.right, value)
	}
	return node
}

func (bst *BST) addNode(value int) {
	bst.node = insertNode(bst.node, value)
}

func (node *Node) display() {
	if node == nil {
		return
	}
	node.left.display()
	fmt.Println("Node value is ", node.value)
	node.right.display()
}

func (bst *BST) display() {
	fmt.Println("BST in-order traversal:\n\n")
	bst.node.display()
}

func searchValue(node *Node, value int) *Node {
	if node == nil {
		return nil
	}
	if value == node.value {
		return node
	}
	if value < node.value {
		return searchValue(node.left, value)
	}
	return searchValue(node.right, value)
}

func (bst *BST) search(value int) *Node {
	return searchValue(bst.node, value)
}

func (bst *BST) height(node *Node) int {

	if node == nil {
		return 0
	}

	leftHeight := bst.height(node.left)
	rightHeight := bst.height(node.right)
	if leftHeight > rightHeight {
		return 1 + leftHeight
	}
	return 1 + rightHeight
}

// Added via AI
func printTree(node *Node, prefix string, label string) {
	if node == nil {
		return
	}
	fmt.Printf("%s%s%d\n", prefix, label, node.value)
	printTree(node.left, prefix+"│   ", "├── L: ")
	printTree(node.right, prefix+"    ", "└── R: ")
}

func (bst *BST) printStructure() {
	fmt.Printf("Root: %d\n", bst.node.value)
	printTree(bst.node.left, "│   ", "├── L: ")
	printTree(bst.node.right, "    ", "└── R: ")
}

// func main() {
// 	bst := &BST{}
// 	bst.addNode(5)
// 	bst.addNode(3)
// 	bst.addNode(7)
// 	bst.addNode(1)
// 	bst.addNode(4)
// 	bst.display()
//
// 	if n := bst.search(3); n != nil {
// 		fmt.Println("Found:", n.value)
// 	} else {
// 		fmt.Println("Not found")
// 	}
// }
