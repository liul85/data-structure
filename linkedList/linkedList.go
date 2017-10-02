package linkedList

import (
	"errors"
	"fmt"
)

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

func (list *LinkedList) Get(i int64) (error, interface{}) {
	err, node := list.getNode(i)

	if err != nil {
		return err, nil
	}

	return nil, node.data
}

func (list *LinkedList) getNode(i int64) (error, *Node) {
	var p int64
	node := list.head

	for {
		if i == p && node != nil {
			return nil, node
		}

		if node == nil || node.next == nil {
			return errors.New("not found"), nil
		}

		node = node.next
		p++
	}
}

func (list *LinkedList) Insert(data interface{}, index int64) error {
	if list.head == nil {
		if index == 0 {
			list.head = &Node{data: data}
			return nil
		}

		return errors.New("can't insert to empty list with index more than 0.")
	}

	err, prevNode := list.getNode(index - 1)

	if err != nil {
		return errors.New("can not insert to list.")
	}

	newNode := Node{data: data}
	prevNode.next, newNode.next = &newNode, prevNode.next

	return nil
}

func (list *LinkedList) Del(index int64) error {
	err, node := list.getNode(index)
	if err != nil {
		return errors.New(fmt.Sprintf("can not get deleted node: %s", err))
	}

	if index == 0 {
		list.head = node.next
		return nil
	}

	e, prevNode := list.getNode(index - 1)

	if e != nil {
		return errors.New(fmt.Sprintf("can not get prev node of deleted: %s", e))
	}

	prevNode.next = node.next

	return nil
}

func (list *LinkedList) Length() int64 {
	var length int64
	node := list.head
	for {
		if node != nil {
			node = node.next
		} else {
			return length
		}
		length++
	}
}

func New(arg ...interface{}) *LinkedList {
	list := LinkedList{}
	for _, n := range arg {
		list.Append(n)
	}
	return &list
}
