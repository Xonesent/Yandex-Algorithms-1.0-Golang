package main

import "fmt"

type node struct {
	value  int
	length int
	left   *node
	right  *node
}

type tree struct {
	root *node
}

func main() {
	tree := &tree{}
	var value int = 1
	values := make(map[int]bool)
	for {
		fmt.Scan(&value)
		if value == 0 {
			break
		}
		if !values[value] {
			values[value] = true
			tree.Insert(value)
		}
	}
}

func initialization(value int, length int) *node {
	return &node{value: value, length: length, left: nil, right: nil}
}

func (root *tree) Insert(value int) {
	if root.root == nil {
		root.root = initialization(value, 1)
		fmt.Print(1, " ")
	} else {
		var length int = 1
		root.root.insert(value, &length)
		fmt.Print(length, " ")
	}
}

func (root *node) insert(value int, length *int) {
	if value >= root.value {
		if root.right == nil {
			*length++
			elem := initialization(value, *length)
			root.right = elem
		} else {
			*length++
			root.right.insert(value, length)
		}
	} else {
		if root.left == nil {
			*length++
			elem := initialization(value, *length)
			root.left = elem
		} else {
			*length++
			root.left.insert(value, length)
		}
	}
}
