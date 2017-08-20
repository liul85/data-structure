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

func TestDeQueueFromEmptyQueue(t *testing.T) {
	queue := New()

	err, _ := queue.Dequeue()

	if err == nil {
		t.Error("Expected error, but did not get")
	}
}

func TestDeQueue(t *testing.T) {
	queue := New()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	err, data := queue.Dequeue()

	if err != nil || data != 1 {
		t.Errorf("Expected data is 1, but got %v\n", data)
	}

	err, data = queue.Dequeue()

	if err != nil || data != 2 {
		t.Errorf("Expected data is 2, but got %v\n", data)
	}

	err, data = queue.Dequeue()

	if err != nil || data != 3 {
		t.Errorf("Expected data is 3, but got %v\n", data)
	}

	err, data = queue.Dequeue()

	if err == nil || data != nil {
		t.Error("Expected error, but did not get")
	}
}

func TestEmptyQueueLength(t *testing.T) {
	queue := New()

	length := queue.Length()

	if length != 0 {
		t.Errorf("Expected 0, but got %v\n", length)
	}
}

func TestQueueLength(t *testing.T) {
	queue := New()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	length := queue.Length()

	if length != 3 {
		t.Errorf("Expected 3, but got %v\n", length)
	}
}
