package queue

import (
	"testing"
)

func TestEmptyQueue(t *testing.T) {
	queue := New()

	if queue.head != nil || queue.tail != nil {
		t.Error("Error creating a new queue.")
	}
}

func TestCreateNewQueue(t *testing.T) {
	queue := New("a", "b", "c")

	if queue.head.data != "a" {
		t.Errorf("Expected head is a, actual is %v\n", queue.head.data)
	}

	if queue.tail.data != "c" {
		t.Errorf("Expected tail is c, actual is %v\n", queue.tail.data)
	}
}

func TestEnqueueToEmptyQueue(t *testing.T) {
	queue := New()

	queue.Enqueue(1)

	if queue.head.data != 1 {
		t.Errorf("Expected head is 1, actual is %v\n", queue.head.data)
	}

	if queue.tail.data != 1 {
		t.Errorf("Expected tail is 1, actual is %v\n", queue.tail.data)
	}

	queue.Enqueue(2)
	queue.Enqueue(3)

	if queue.head.data != 1 {
		t.Errorf("Expected head is 1, actual is %v\n", queue.head.data)
	}

	if queue.head.next.data != 2 {
		t.Errorf("Expected head is 2, actual is %v\n", queue.head.data)
	}

	if queue.tail.data != 3 {
		t.Errorf("Expected tail is 3, actual is %v\n", queue.tail.data)
	}
}
