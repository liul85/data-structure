package linkedList

import (
	"testing"
)

func TestEmptyLinkedList(t *testing.T) {
	list := LinkedList{}

	if list.Head != nil {
		t.Errorf("expect nil, acturl is %v", list.Head)
	}
}

func TestAppendOneElemToEmptyLinkedList(t *testing.T) {
	list := LinkedList{}

	list.Append(1)

	if list.Head.Data != 1 {
		t.Errorf("expect data equals 1, actual is %v", list.Head.Data)
	}
}

func TestAppendOneElemToNoneEmptyLinkedList(t *testing.T) {
	list := LinkedList{}
	list.Append(1)
	list.Append(2)

	if list.Head.Next.Data != 2 {
		t.Errorf("expect data equals 2, actual is %v", list.Head.Next.Data)
	}
}

func TestGetElemFromList(t *testing.T) {
	list := LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	err, value := list.Get(4)

	if value != 4 || err != nil {
		t.Errorf("expected 4, but got %v\n", value)
	}
}

func TestGetElemFromListExceedsTheLengh(t *testing.T) {
	list := LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	err, value := list.Get(5)

	if value != nil || err == nil {
		t.Errorf("expected 4, but got %v\n", value)
	}
}

func TestGetElemFromEmptyList(t *testing.T) {
	list := LinkedList{}

	err, value := list.Get(1)

	if value != nil || err == nil {
		t.Errorf("expected nil, but got %v\n", value)
	}
}
