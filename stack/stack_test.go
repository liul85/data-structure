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

func TestNewStack(t *testing.T) {
	stack := New(2, 3, 6, 9)

	if stack.top.data != 9 {
		t.Error("Failed to new stack")
	}

	if stack.Empty() != false {
		t.Error("Failed to create new stack")
	}
}

func TestStackEmpty(t *testing.T) {
	stack := New()

	if stack.Empty() != true {
		t.Error("Expected empty stack, actual is not empty.")
	}
}

func TestGetTopData(t *testing.T) {
	stack := New(2, 3)
	top := stack.GetTop()

	if top != 3 {
		t.Errorf("Expected 3, but got %v\n", top)
	}
}

func TestGetTopDataFromEmptyStack(t *testing.T) {
	stack := New()

	if stack.GetTop() != nil {
		t.Error("Expected nil, but actual is not nil")
	}
}

func TestEmptyStackLength(t *testing.T) {
	stack := New()
	length := stack.Length()

	if length != 0 {
		t.Errorf("Expected 0, but actual is %v\n", length)
	}
}

func TestStackLength(t *testing.T) {
	stack := New(1, 2, 3, 4)
	length := stack.Length()

	if length != 4 {
		t.Errorf("Expected 4, but actual is %v\n", length)
	}
}
