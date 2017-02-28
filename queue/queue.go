package queue

import (
	"github.com/cjang5/ds/stack"
)

// Queue - A queue implementation made up of Items
// and uses two stacks
type Queue struct {
	length int
	s1, s2 *stack.Stack
}

// New creates a new Queue (a pointer to one)
// and initializes length and s1/s2
func New() *Queue {
	q := new(Queue)
	q.length = 0
	q.s1 = stack.New()
	q.s2 = stack.New()

	return q
}

// Peek returns the value at the front of the Queue
// and returns type interface{} since the Queue's item type may vary
func (q *Queue) Peek() interface{} {
	// check if Queue is empty
	if q.IsEmpty() {
		return nil
	}

	// Otherwise return the head element (not an item, but interface{})
	if !q.s2.IsEmpty() {
		return q.s2.Peek()
	}
	q.transfer()
	return q.s2.Peek()
}

// transfer moves everything from the s1 stack to the s2 stack, so that
// the popping order in s2 is correct
func (q *Queue) transfer() {
	for q.s1.Peek() != nil {
		q.s2.Push(q.s1.Pop())
	}
}

// Enqueue puts a new item at the back of the Queue
func (q *Queue) Enqueue(newVal interface{}) {
	q.s1.Push(newVal)
}

// Dequeue removes the item at the front of the Queue and returns it
func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	if !q.s2.IsEmpty() {
		q.length--
		return q.s2.Pop()
	}

	q.transfer()
	return q.Dequeue()
}

// IsEmpty returns whether or not the Queue is empty
func (q *Queue) IsEmpty() bool {
	return q.s1.IsEmpty() && q.s2.IsEmpty()
}
