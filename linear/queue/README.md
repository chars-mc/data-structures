## Queue

Is a list of elements of type **FIFO (First-In-First-Out)** and it has two parts, the **REAR** and **FRONT**.

### Operations

- **enqueue**: inserts data on `REAR`. `o(1)`
- **dequeue**: removes data from the `FRONT`. `o(1)`
- **peek**: returns an specific element. `o(n)`
- **isFull**: if the queue is full, throws `overflow` condition.
- **isEmpty**: if the queue is empty, throws `underflow` condition.

### Types

- **Linear**: insertion and deletions occurs on one end, on the `REAR` and
  `FRONT` respectively.
- **Circular**: the last element is connected to the first element.
- **Priority**: each element has a priotity, so the elements are sorted by their
  priority, and the insertion is performed as an element is added, and deletion
  by priority.
- **Deque**: insertion and deletion can occur from the `REAR` or from the
  `FRONT`, that means it doesn't follow the FIFO principle.
