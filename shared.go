package queue

import "errors"

var (
	// Queues will panic with this error when empty and
	// the Dequeue method is called.
	ErrEmptyQueue = errors.New("queue: the queue is empty and the requested operation could not be performed")

	// Queues will panic with this error when they encounter a
	// value in the underlying list that isn't of the expected
	// type. This SHOULD NOT ever happen, and if it does it
	// indicates that the underlying `container/list` was
	// exported and manipulated by outside code, or that there
	// is a bug in this code. Both are bad and shouldn't be
	// allowed to happen!
	ErrInvalidType = errors.New("queue: invalid type encountered - this indicates a bug.")
)
