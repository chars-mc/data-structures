package main

import (
	"fmt"

	"github.com/chars-mc/data-structures/linear/models"
)

// linkedList represents a list of Nodes
type linkedList struct {
	head   *models.Node
	length int
}

// prepend adds a Node to the list
func (l *linkedList) prepend(n *models.Node) {
	previousNode := l.head     // get the last node
	l.head = n                 // set n as the head of the list
	l.head.Next = previousNode // set n.Next to the previous node
	l.length++
}

// printList prints the data and pointer of all the nodes on the list
func (l linkedList) printList() {
	head := l.head
	for head.Next != nil {
		fmt.Printf("[%s | %p] -> ", head.Data, head.Next)
		head = head.Next
	}
	// print the last element on the list
	fmt.Printf("[%s | %p] -> nil\n", head.Data, head.Next)
}

// deleteWithValue deletes the first element that matches the parameter 'v'
func (l *linkedList) deleteWithValue(v string) {
	// if the list is empty, return
	if l.length == 0 {
		return
	}

	// if the list head has the same value as v, change the head to the next Node
	if l.head.Data == v {
		l.head = l.head.Next
		l.length--
		return
	}

	previousNode := l.head
	for previousNode.Next.Data != v {
		// if v is not in the list, return
		if previousNode.Next.Next == nil {
			return
		}
		previousNode = previousNode.Next
	}

	previousNode.Next = previousNode.Next.Next
	l.length--
}

func main() {
	l := linkedList{}
	n := &models.Node{Data: "hello"}
	n2 := &models.Node{Data: "world"}
	n3 := &models.Node{Data: "!!!"}
	n4 := &models.Node{Data: "!!!"}

	l.prepend(n)
	l.prepend(n2)
	l.prepend(n3)
	l.prepend(n4)

	l.printList()
	l.deleteWithValue("!!!")
	l.deleteWithValue("v")
	l.printList()
}
