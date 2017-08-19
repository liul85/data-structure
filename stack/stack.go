package stack

type Node struct {
	data interface{}
	next *Node
}

type Stack struct {
	top *Node
}

func (stack *Stack) push(data interface{}) {
	newNode := Node{data: data}
	newNode.next = stack.top
	stack.top = &newNode
}

func (stack *Stack) pop() (interface{}, interface{}) {
	if stack.top == nil {
		return "Empty stack", nil
	}
	data := stack.top.data
	stack.top = stack.top.next
	return nil, data
}
