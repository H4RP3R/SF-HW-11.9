package queue

import (
	"container/list"
	"fmt"
)

var ErrPollFromEmptyQueue = fmt.Errorf("poll from empty queue")

type Queue[T comparable] struct {
	storage *list.List
}

// Size returns the current elements count in queue
func (q Queue[T]) Size() int {
	return q.storage.Len()
}

// Add puts the element at the end of the queue.
func (q *Queue[T]) Add(el T) {
	q.storage.PushBack(el)
}

// Poll retrieves and removes the head of this queue, or returns nil if this queue is empty.
func (q *Queue[T]) Poll() (T, error) {
	var val T
	if q.Size() == 0 {
		return val, ErrPollFromEmptyQueue
	}

	el := q.storage.Front()
	val = el.Value.(T)
	q.storage.Remove(el)
	return val, nil
}

// Remove removes the first occurrence of the element, returns true if there were an occurrence
// otherwise return false.
func (q *Queue[T]) Remove(el T) bool {
	for e := q.storage.Front(); e != nil; e = e.Next() {
		if e.Value == el {
			q.storage.Remove(e)
			return true
		}
	}
	return false
}

func NewQueue[T comparable]() Queue[T] {
	q := Queue[T]{}
	q.storage = list.New()
	return q
}
