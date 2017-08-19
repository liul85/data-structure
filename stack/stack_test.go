package stack

import (
	"testing"
)

func TestPushToEmtpyStack(t *testing.T) {
	stack := Stack{}
	stack.push(1)

	topData := stack.top.data

	if topData != 1 {
		t.Errorf("Expected 1, but got %v\n", topData)
	}
}

func TestPushToStack(t *testing.T) {
	stack := Stack{}
	stack.push(1)
	stack.push(2)
	stack.push(3)

	topData := stack.top.data
	if topData != 3 {
		t.Errorf("Expected 3, got %s\n", topData)
	}
}

func TestPopFromEmptyStack(t *testing.T) {
	stack := Stack{}

	err, data := stack.pop()

	if err == nil || data != nil {
		t.Error("Expected got error, but got nil")
	}
}

func TestPopFromStack(t *testing.T) {
	stack := Stack{}
	stack.push(1)

	err, data := stack.pop()

	if err != nil || data != 1 {
		t.Errorf("Expected 1, but got %v\n", data)
	}

	err, data = stack.pop()

	if err == nil || data != nil {
		t.Error("Expected got error, but got nil")
	}
}

func TestPopMoreFromStack(t *testing.T) {
	stack := Stack{}
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)

	stack.pop()
	stack.pop()
	_, data3 := stack.pop()

	if data3 != 2 {
		t.Errorf("Expected 2, but got %v", data3)
	}
}
