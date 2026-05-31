package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(data string) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (l *LinkedList) Display() {
	current := l.head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}

func main() {
	list := &LinkedList{}
	fmt.Print("Creating a single linked list and inserting elements...\n")
	fmt.Println("Enter the number of elements in the linked list")
	var n int
	fmt.Scanln(&n)
	fmt.Println("Enter the elements of the linked list")
	for i := 0; i < n; i++ {
		var element string
		fmt.Scanln(&element)
		list.Insert(element)
		list.Display()
	}
}
