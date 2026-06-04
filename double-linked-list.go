package main

import "fmt"

type Node struct {
	value int
	prev  *Node
	next  *Node
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
}

func (dll *DoubleLinkedList) Insert(value int) {
	newNode := &Node{value: value}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
	}
}

func (dll *DoubleLinkedList) Display() {
	current := dll.head
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Println()
}

func main() {
	dll := &DoubleLinkedList{}
	fmt.Println("Enter the number of elements to insert:")
	var n int
	fmt.Scanln(&n)
	fmt.Println("Enter the elements:")
	for i := 0; i < n; i++ {
		var value int
		fmt.Scanln(&value)
		dll.Insert(value)
		dll.Display()
	}
}
