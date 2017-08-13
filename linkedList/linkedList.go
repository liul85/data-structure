package linkedList

type Node struct {
	next *Node
	data interface{}
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Append(data interface{}) {
	node := Node{data: data}

	if list.head == nil {
		list.head = &node
	} else {
		list.getLast().next = &node
	}
}

func (list *LinkedList) getLast() *Node {
	node := list.head
	for {
		if node.next == nil {
			return node
		}
		node = node.next
	}
}

func (list *LinkedList) Get(i int) (interface{}, interface{}) {
	p := 1
	node := list.head

	for {
		if i == p && node != nil {
			return nil, node.data
		}

		if node == nil || node.next == nil {
			return "not found", nil
		}

		node = node.next
		p++
	}
}
