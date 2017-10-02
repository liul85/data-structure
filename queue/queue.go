package queue

import (
	"errors"
)

type node struct {
	data interface{}
	next *node
}

type queue struct {
	head *node
	tail *node
}

func New(args ...interface{}) *queue {
	q := &queue{}
	for _, v := range args {
		q.Enqueue(v)
	}

	return q
}

func (q *queue) Enqueue(data interface{}) {
	node := &node{data: data}
	if q.IsEmpty() {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
}

func (q *queue) Dequeue() (error, interface{}) {
	if q.IsEmpty() {
		return errors.New("empty queue"), nil
	}

	data := q.head.data

	if q.head == q.tail {
		q.tail = nil
	}

	q.head = q.head.next
	return nil, data
}

func (q *queue) IsEmpty() bool {
	return q.head == nil && q.tail == nil
}

func (q *queue) Length() int64 {
	if q.IsEmpty() {
		return 0
	}

	var length int64 = 1
	node := q.head
	for {
		if node.next != nil {
			length++
			node = node.next
		} else {
			return length
		}
	}
}
