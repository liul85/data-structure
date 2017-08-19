package stack

import (
	"testing"
)

func TestPushToEmtpyStack(t *testing.T) {
	stack := Stack{}
	stack.Push(1)

	topData := stack.top.data

	if topData != 1 {
		t.Errorf("Expected 1, but got %v\n", topData)
	}
}

func TestPushToStack(t *testing.T) {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	topData := stack.top.data
	if topData != 3 {
		t.Errorf("Expected 3, got %s\n", topData)
	}
}

func TestPopFromEmptyStack(t *testing.T) {
	stack := Stack{}

	err, data := stack.Pop()

	if err == nil || data != nil {
		t.Error("Expected got error, but got nil")
	}
}

func TestPopFromStack(t *testing.T) {
	stack := Stack{}
	stack.Push(1)

	err, data := stack.Pop()

	if err != nil || data != 1 {
		t.Errorf("Expected 1, but got %v\n", data)
	}

	err, data = stack.Pop()

	if err == nil || data != nil {
		t.Error("Expected got error, but got nil")
	}
}

func TestPopMoreFromStack(t *testing.T) {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	stack.Pop()
	stack.Pop()
	_, data3 := stack.Pop()

	if data3 != 2 {
		t.Errorf("Expected 2, but got %v", data3)
	}
}
