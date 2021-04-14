package main

import "fmt"

// stack represents a stack that hold a slice of strings
type stack struct {
	items []string
}

// push adds an element
func (s *stack) push(el string) {
	s.items = append(s.items, el)
}

// isEmpty check if the stack is empty and return a boolean value
func (s stack) isEmpty() bool {
	return len(s.items) == 0
}

// pop removes an element
func (s *stack) pop() {
	s.items = s.items[:len(s.items)-1]
}

// count returns the lenght of the stack
func (s stack) count() int {
	return len(s.items)
}

// display prints the stack elements
func (s stack) display() {
	fmt.Println(s)
}

func main() {
	s := stack{}
	s.push("hello")
	s.push("gophers")
	s.push("!!!")

	s.display()
	fmt.Printf("Stack lenght: %d\n", s.count())
	fmt.Printf("Stack is empty? %t\n", s.isEmpty())

	s.pop()
	s.display()
}
