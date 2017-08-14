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

func (list *LinkedList) Get(i int64) (interface{}, interface{}) {
	err, node := list.getNode(i)

	if err != nil {
		return err, nil
	}

	return nil, node.data
}

func (list *LinkedList) getNode(i int64) (interface{}, *Node) {
	var p int64
	node := list.head

	for {
		if i == p && node != nil {
			return nil, node
		}

		if node == nil || node.next == nil {
			return "not found", nil
		}

		node = node.next
		p++
	}
}

func (list *LinkedList) Insert(data interface{}, index int64) interface{} {
	if list.head == nil {
		if index == 0 {
			list.head = &Node{data: data}
			return nil
		}

		return "can't insert to empty list with index more than 0."
	}

	err, prevNode := list.getNode(index - 1)

	if err != nil {
		return "can not insert to list."
	}

	newNode := Node{data: data}
	newNode.next = prevNode.next
	prevNode.next = &newNode

	return nil
}
