package linkedList

import (
	"testing"
)

func TestEmptyLinkedList(t *testing.T) {
	list := LinkedList{}

	if list.head != nil {
		t.Errorf("expect nil, acturl is %v", list.head)
	}
}

func TestAppendOneElemToEmptyLinkedList(t *testing.T) {
	list := LinkedList{}

	list.Append(1)

	if list.head.data != 1 {
		t.Errorf("expect data equals 1, actual is %v", list.head.data)
	}
}

func TestAppendOneElemToNoneEmptyLinkedList(t *testing.T) {
	list := LinkedList{}
	list.Append(1)
	list.Append(2)

	if list.head.next.data != 2 {
		t.Errorf("expect data equals 2, actual is %v", list.head.next.data)
	}
}

func TestGetElemFromList(t *testing.T) {
	list := LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	err, value := list.Get(3)

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

func TestInsertElemInEmptyList(t *testing.T) {
	list := LinkedList{}

	list.Insert(1, 0)

	err, value := list.Get(0)

	if value != 1 || err != nil {
		t.Errorf("expected 1, but got %v\n", value)
	}
}

func TestInsertElemInEmptyListWithErr(t *testing.T) {
	list := LinkedList{}

	err := list.Insert(1, 9)

	if err == nil {
		t.Error("expected got error, but got nil\n")
	}
}

func TestInsertElemInNoneEmptyList(t *testing.T) {
	list := LinkedList{}

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	list.Insert(9, 2)

	var err, value interface{}

	err, value = list.Get(2)

	if err != nil || value != 9 {
		t.Errorf("expected 9, but got %v\n", value)
	}

	err, value = list.Get(3)

	if err != nil || value != 3 {
		t.Errorf("expected 3, but got %v\n", value)
	}
}

func TestInsertElemInNonEmptyListWithErr(t *testing.T) {
	list := LinkedList{}

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	err := list.Insert(88, 9)

	if err == nil {
		t.Errorf("expected err, but got nil")
	}
}

func TestInsertElemToTailOfLink(t *testing.T) {
	list := LinkedList{}

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	list.Insert(5, 4)

	err, data := list.Get(4)

	if err != nil || data != 5 {
		t.Errorf("expected 5, but got %v\n", data)
	}
}

func TestDelElemInEmptyList(t *testing.T) {
	list := LinkedList{}

	err := list.Del(3)

	if err == nil {
		t.Error("expected err, but got nil")
	}
}

func TestDelHeadWhenLengthIsOne(t *testing.T) {
	list := LinkedList{}

	list.Append(1)

	err := list.Del(0)

	if err != nil {
		t.Errorf("expected success, got err: %v", err)
	}

	if e, data := list.Get(0); e == nil || data != nil {
		t.Error("not empty list after delete head.")
	}
}

func TestDelHeadWhenLengthMoreThanOne(t *testing.T) {
	list := LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	err := list.Del(0)

	if err != nil {
		t.Errorf("Error happend when delete head: %v\n", err)
	}

	e, data := list.Get(0)

	if e != nil || data != 2 {
		t.Errorf("Expected 2, but got %v\n", data)
	}

}

func TestDelElem(t *testing.T) {
	list := LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	err := list.Del(2)

	if err != nil {
		t.Errorf("Error happend when delete head: %v\n", err)
	}

	e, data := list.Get(2)

	if e != nil || data != 4 {
		t.Errorf("Expected 4, but got %v\n", data)
	}
}

func TestDelTail(t *testing.T) {
	list := LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	err := list.Del(2)

	if err != nil {
		t.Errorf("Error happend when delete head: %v\n", err)
	}

	e, data := list.Get(2)

	if e == nil || data != nil {
		t.Errorf("Expected nil, but got %v\n", data)
	}
}
