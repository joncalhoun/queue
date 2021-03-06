package queue

import (
	"container/list"
)

func NewIntSlice() *IntSlice {
	return &IntSlice{list.New()}
}

// IntSlice is a queue implementation for the []int type.
// Behind the scenes it is a linked list FIFO queue
// that uses container/list under the hood. The primary
// motivation in creating this type is to allow the compiler
// to verify that we are using the correct types with our
// queue rather than dealing with the interface{} type in
// the rest of our code.
type IntSlice struct {
	list *list.List
}

// Len returns the total length of the queue
func (q *IntSlice) Len() int {
	return q.list.Len()
}

// Enqueue adds an item to the back of the queue
func (q *IntSlice) Enqueue(i []int) {
	q.list.PushBack(i)
}

// Dequeue removes and returns the front item in the queue
func (q *IntSlice) Dequeue() []int {
	if q.list.Len() == 0 {
		// You could opt to return errors here, but I personally
		// prefer to leave length checking up to end users kinda
		// like bounds checking in slices.
		panic(ErrEmptyQueue)
	}

	raw := q.list.Remove(q.list.Front())
	if typed, ok := raw.([]int); ok {
		return typed
	}

	// This won't ever happen unless someone has access to
	// insert things into the list with an invalid type or
	// your code has bug.
	panic(ErrInvalidType)
}
