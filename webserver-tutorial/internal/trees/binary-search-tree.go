package trees

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

func (bst *BST) AddNode(value int) {
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

func (bst *BST) Display() {
	fmt.Println("BST in-order traversal")
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

func (bst *BST) Search(value int) *Node {
	return searchValue(bst.node, value)
}

func collectInOrder(node *Node, result *[]int) {
	if node == nil {
		return
	}
	collectInOrder(node.left, result)
	*result = append(*result, node.value)
	collectInOrder(node.right, result)
}

func (bst *BST) InOrder() []int {
	result := []int{}
	collectInOrder(bst.node, &result)
	return result
}

func (bst *BST) Height(node *Node) int {

	if node == nil {
		return 0
	}

	leftHeight := bst.Height(node.left)
	rightHeight := bst.Height(node.right)
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

func (bst *BST) PrintStructure() {
	fmt.Printf("Root: %d\n", bst.node.value)
	printTree(bst.node.left, "│   ", "├── L: ")
	printTree(bst.node.right, "    ", "└── R: ")
}
