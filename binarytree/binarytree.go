package binarytree

import (
	"fmt"
)

type binaryTree struct {
	data   interface{}
	lchild *binaryTree
	rchild *binaryTree
}

func createTree(tree *binaryTree, data *[]interface{}) {
	if len(*data) == 0 {
		return
	}

	v := (*data)[:1][0]
	*data = (*data)[1:]

	if fmt.Sprintf("%s", v) == "#" {
		return
	}

	tree.data = v

	if tree.lchild == nil {
		tree.lchild = &binaryTree{}
		createTree(tree.lchild, data)
	}

	if tree.rchild == nil {
		tree.rchild = &binaryTree{}
		createTree(tree.rchild, data)
	}
}

func PreOrderNew(data []interface{}) *binaryTree {
	tree := &binaryTree{}
	for _, v := range data {
		if tree.data == nil {
			createBiTree(tree, v)
		} else if tree.lchild == nil {
			lchild := &binaryTree{}
			tree.lchild = lchild
			createBiTree(tree.lchild, v)
		} else if tree.rchild == nil {
			rchild := &binaryTree{}
			tree.rchild = rchild
			createBiTree(tree.rchild, v)
		}
	}
	return tree
}

func createBiTree(tree *binaryTree, data interface{}) {
	if data == "#" {
		tree.data = nil
	} else {
		tree.data = data
	}
}

func PreOrderTraverse(tree binaryTree) []interface{} {
	if tree.data == nil {
		return nil
	}

	result := make([]interface{}, 1)

	result[0] = tree.data

	if tree.lchild != nil {
		for _, v := range PreOrderTraverse(*tree.lchild) {
			result = append(result, v)
		}
	}

	if tree.rchild != nil {
		for _, v := range PreOrderTraverse(*tree.rchild) {
			result = append(result, v)
		}
	}

	return result
}
