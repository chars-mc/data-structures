package models

// Node represent an element that stores a data and a pointer to a Node
type Node struct {
	Data string
	Next *Node
}
