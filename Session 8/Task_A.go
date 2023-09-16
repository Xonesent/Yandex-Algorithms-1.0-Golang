package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}

type tree struct {
	length int
	root   *node
}

func main() {
	tree := &tree{length: 0}
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
	fmt.Println(tree.length)
}

func initialization(value int) *node {
	return &node{value: value, left: nil, right: nil}
}

func (root *tree) Insert(value int) {
	if root.root == nil {
		root.root = initialization(value)
		root.length = 1
	} else {
		var length int = 1
		root.root.insert(value, &length)
		if length > root.length {
			root.length = length
		}
	}
}

func (root *node) insert(value int, length *int) {
	elem := initialization(value)
	if value >= root.value {
		if root.right == nil {
			root.right = elem
			*length++
		} else {
			*length++
			root.right.insert(value, length)
		}
	} else {
		if root.left == nil {
			root.left = elem
			*length++
		} else {
			*length++
			root.left.insert(value, length)
		}
	}
}
