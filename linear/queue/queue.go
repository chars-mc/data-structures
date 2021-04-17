package main

import "fmt"

// queue represents a queue that holds a slice of strings
type queue struct {
	items []string
}

// enqueue adds a value to the 'front'
func (q *queue) enqueue(s string) {
	q.items = append(q.items, s)
}

// dequeue removes an item from the 'rear'
// and RETURNs the removed value
func (q *queue) dequeue() string {
	toDelete := q.items[0]
	q.items = q.items[1:len(q.items)]
	return toDelete
}

func main() {
	q := queue{}

	q.enqueue("hello")
	q.enqueue("gophers")
	q.enqueue("!!!")

	deleted := q.dequeue()
	fmt.Printf("value deleted: %s\n", deleted)
	fmt.Println(q)
}
