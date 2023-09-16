package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}

type tree struct {
	root *node
}

func main() {
	tree := &tree{}
	var value int
	for {
		fmt.Scan(&value)
		if value == 0 {
			break
		}
		tree.Insert(value)
	}
	ans := tree.Sorted()

	for _, value := range ans {
		fmt.Println(value)
	}
}

func initialization(value int) *node {
	return &node{value: value, left: nil, right: nil}
}

func (root *tree) Insert(value int) {
	if root.root == nil {
		root.root = initialization(value)
	} else {
		root.root.insert(value)
	}
}

func (root *node) insert(value int) {
	if value > root.value {
		if root.right == nil {
			elem := initialization(value)
			root.right = elem
		} else {
			root.right.insert(value)
		}
	} else if value < root.value {
		if root.left == nil {
			elem := initialization(value)
			root.left = elem
		} else {
			root.left.insert(value)
		}
	}
}

func (root *tree) Sorted() []int {
	ans := make([]int, 0)
	if root.root != nil {
		root.root.sorted(&ans)
	}
	return ans
}

func (root *node) sorted(ans *[]int) {
	if root != nil {
		root.left.sorted(ans)
		if root.left != nil && root.right != nil {
			*ans = append(*ans, root.value)
		}
		root.right.sorted(ans)
	}
}
