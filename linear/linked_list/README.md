## Linked list

A linked list is a sequential collection of elements (nodes), and this elements are linked to each other.

A node consists of the data and a pointer to the next node:

```go
type Node struct {
	Data string
	Next *Node
}
```

And the list:

```go
type LinkedList struct {
	head   *Node
	length int
}
```

### Advantages

- Dynamic length and memory efficient.
- Easy insertion on deletion at the beggining of the list (head). `O(1)`.

### Disdvantages

- More memory usage cause it stores two variables, the data and the pointer.
- Traversal: difficult access to a specific element. `O(n)`.
- Reverse traversing: in a double-linked list requires more memory usage
