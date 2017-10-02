package binarytree

import (
	"testing"
)

func TestPreOrder(t *testing.T) {
	data := []interface{}{"A"}
	tree := PreOrderNew(data)

	result := PreOrderTraverse(*tree)

	if result[0] != "A" {
		t.Errorf("Expected A, but got %v\n", result[0])
	}
}

func TestPreOrderWithLChild(t *testing.T) {
	data := []interface{}{"A", "B"}
	tree := PreOrderNew(data)

	result := PreOrderTraverse(*tree)

	if len(result) != 2 {
		t.Errorf("Expected length of result 2, but got %v", len(result))
	}

	if result[0] != "A" {
		t.Errorf("Expected root node is A, but got %v", result[0])
	}

	if result[1] != "B" {
		t.Errorf("Expected left child is B, but got %v", result[1])
	}
}

func TestPreOrderWithLRchild(t *testing.T) {
	data := []interface{}{"A", "B", "C"}
	tree := PreOrderNew(data)

	result := PreOrderTraverse(*tree)

	if len(result) != 3 {
		t.Errorf("Expected length of result 3, but got %v", len(result))
	}

	if result[0] != "A" || result[1] != "B" || result[2] != "C" {
		t.Errorf("got wrong traverse data")
	}
}
