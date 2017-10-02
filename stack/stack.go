package stack

import (
	"errors"
)

type Node struct {
	data interface{}
	next *Node
}

type Stack struct {
	top *Node
}

func (stack *Stack) Push(data interface{}) {
	newNode := Node{data: data}
	newNode.next = stack.top
	stack.top = &newNode
}

func (stack *Stack) Pop() (error, interface{}) {
	if stack.top == nil {
		return errors.New("Empty stack"), nil
	}
	data := stack.top.data
	stack.top = stack.top.next
	return nil, data
}

func New(args ...interface{}) *Stack {
	stack := Stack{}
	for _, n := range args {
		stack.Push(n)
	}
	return &stack
}

func (stack *Stack) Empty() bool {
	return stack.top == nil
}

func (stack *Stack) GetTop() interface{} {
	if stack.Empty() {
		return nil
	}

	return stack.top.data
}

func (stack *Stack) Length() int64 {
	if stack.Empty() {
		return 0
	}

	var length int64 = 1
	node := stack.top

	for {
		if node.next != nil {
			node = node.next
			length++
		} else {
			return length
		}
	}
}
