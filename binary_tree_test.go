package main

import (
	"fmt"
	"testing"
)

type node struct {
	value  int
	leftN  *node
	rightN *node
}

func (n *node) insert(value int) {
	if value == n.value {
		return
	}

	if n.value > value {
		if n.leftN == nil {
			n.leftN = &node{value: value}
		} else {
			n.leftN.insert(value)
		}
	} else {
		if n.rightN == nil {
			n.rightN = &node{value: value}
		} else {
			n.rightN.insert(value)
		}
	}
}

func (n *node) inOrderTraversal() []int {
	res := make([]int, 0)
	if n.leftN != nil {
		res = append(res, n.leftN.inOrderTraversal()...)
	}
	res = append(res, n.value)
	if n.rightN != nil {
		res = append(res, n.rightN.inOrderTraversal()...)
	}
	return res
}

func (n *node) preOrderTraversal() []int {
	res := make([]int, 0)
	res = append(res, n.value)
	if n.leftN != nil {
		res = append(res, n.leftN.preOrderTraversal()...)
	}
	if n.rightN != nil {
		res = append(res, n.rightN.preOrderTraversal()...)
	}
	return res
}

func buildTree(arr []int) *node {
	root := &node{value: arr[0]}
	for _, a := range arr {
		root.insert(a)
	}
	return root
}

func (n *node) print() {
	fmt.Println(n.value)
	if n.leftN != nil {
		n.leftN.print()
	}
	if n.rightN != nil {
		n.rightN.print()
	}
}

func (n *node) findMin() int {
	if n.leftN != nil {
		return n.leftN.findMin()
	} else {
		return n.value
	}
}

func (n *node) findMax() int {
	if n.rightN != nil {
		return n.rightN.findMax()
	} else {
		return n.value
	}
}

func (n *node) sum() int {
	sum := n.value
	if n.leftN != nil {
		sum += n.leftN.sum()
	}
	if n.rightN != nil {
		sum += n.rightN.sum()
	}
	return sum
}

func (n *node) delete(root *node, value int) *node {
	if root == nil {
		return root
	}

	if value < root.value {
		root.leftN = n.delete(root.leftN, value)
	} else if value > root.value {
		root.rightN = n.delete(root.rightN, value)
	} else {
		if root.leftN == nil {
			return root.rightN
		} else {
			temp := root.leftN.findMax()
			root.value = temp
			root.leftN = n.delete(root.leftN, value)
		}
	}

	return root
}

func TestBinaryTree(t *testing.T) {
	arr := []int{17, 4, 1, 20, 9, 23, 18, 34, 18, 4}
	tree := buildTree(arr)
	fmt.Printf("tree.preOrderTraversal(): %v\n", tree.preOrderTraversal())
	fmt.Printf("tree.inOrderTraversal(): %v\n", tree.inOrderTraversal())
}
