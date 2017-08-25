package queue

import (
	"sync"
)

type CircularQueue struct {
	sync.RWMutex
	size     int
	elements []int
	front    int
	rear     int
}

func New(size int) *CircularQueue {
	return &CircularQueue{
		size:     size,
		elements: make([]int, size),
		front:    -1,
		rear:     -1,
	}
}

func (q *CircularQueue) IsFull() (bool) {
	q.RLock()
	defer q.RUnlock()
	if (q.rear == q.size-1 && q.front == 0) || q.rear == q.front-1 {
		return true
	}
	return false
}

func (q *CircularQueue) IsEmpty() bool {
	q.RLock()
	defer q.RUnlock()
	return q.front == -1
}

func (q *CircularQueue) EnQueue(element int) {
	q.RLock()
	defer q.RUnlock()
	if q.IsFull() {
		panic("queue was full")
	} else if q.front == -1 {
		q.front, q.rear = 0, 0
	} else if q.rear == q.size-1 && q.front != 0 {
		q.rear = 0
	} else {
		q.rear++
	}
	q.elements[q.rear] = element
}

func (q *CircularQueue) DeQueue() (interface{}) {
	q.RLock()
	defer q.RUnlock()
	if q.IsEmpty() {
		return nil
	}
	data := q.elements[q.front]
	q.elements[q.front] = 0
	if q.front == q.rear {
		q.front, q.rear = -1, -1
	} else if q.front == q.size-1 {
		q.front = 0
	} else {
		q.front ++
	}
	return data
}
