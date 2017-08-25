package queue

import (
	"testing"
)

func TestCircularQueueInit(t *testing.T) {
	q := New(3)
	if q.size != 3 {
		t.Errorf("Init queue size was incorrect, got: %d, want: %d.", q.size, 3)
	}
}

func TestCircularQueueEnQueue(t *testing.T) {
	q := New(3)
	q.EnQueue(10)
	if q.elements[0] != 10 {
		t.Errorf("EnQueue was incorrect, first element got: %d, want: %d.", q.elements[0], 10)
	}
	q.EnQueue(20)
	if q.elements[0] != 10 && q.elements[1] != 20 {
		t.Errorf("EnQueue was incorrect, first element got: %d, want: %d.", q.elements[1], 20)
	}
	q.EnQueue(30)
	if q.elements[0] != 10 && q.elements[1] != 20 && q.elements[2] != 30 {
		t.Errorf("EnQueue was incorrect, first element got: %d, want: %d.", q.elements[2], 30)
	}
	func() {
		defer func() { recover() }()
		q.EnQueue(10)
	}()
}

func TestCircularQueueDeQueue(t *testing.T) {
	q := New(2)
	q.EnQueue(10)
	q.EnQueue(20)
	if r := q.DeQueue(); r != 10 {
		t.Errorf("DeQueue was incorrect.")
	}
	if r := q.DeQueue(); r != 20 {
		t.Errorf("DeQueue was incorrect.")
	}
	if r := q.DeQueue(); r != nil {
		t.Errorf("DeQueue was incorrect, can not dequeue when empty queue, result got: %d, want: %#v", r, nil)
	}
}

func TestCircularQueueIsFull(t *testing.T) {
	q := New(1)
	q.EnQueue(10)
	if !q.IsFull() {
		t.Errorf("IsFull was incorrect.")
	}
}

func TestCircularQueueIsEmpty(t *testing.T) {
	q := New(1)
	if !q.IsEmpty() {
		t.Errorf("IsEmpty was incorrect.")
	}
}
