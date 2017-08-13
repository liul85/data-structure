package linkedList

type Node struct {
	Next *Node
	Data interface{}
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Append(data interface{}) {
	node := Node{Data: data}

	if list.Head == nil {
		list.Head = &node
	} else {
		list.getLast().Next = &node
	}
}

func (list *LinkedList) getLast() *Node {
	node := list.Head
	for {
		if node.Next == nil {
			return node
		}
		node = node.Next
	}
}

func (list *LinkedList) Get(i int) (interface{}, interface{}) {
	p := 1
	node := list.Head

	for {
		if i == p && node != nil {
			return nil, node.Data
		}

		if node == nil || node.Next == nil {
			return "not found", nil
		}

		node = node.Next
		p++
	}
}
